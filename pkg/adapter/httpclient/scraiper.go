package main

import (
	"bufio"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql" // Register MySQL Driver
	"github.com/jmoiron/sqlx"
	"github.com/smockoro/scrape-race/pkg/adapter/repository"
	"github.com/smockoro/scrape-race/pkg/domain/model"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func main() {
	url := "https://db.netkeiba.com/race/201806040511/"

	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	utfBody := transform.NewReader(bufio.NewReader(res.Body), japanese.EUCJP.NewDecoder())

	doc, err := goquery.NewDocumentFromReader(utfBody)
	if err != nil {
		fmt.Println(err)
	}
	resultTable := doc.Find("table.race_table_01")
	if err != nil {
		fmt.Println(err)
	}

	JockeySlice := make([]*model.Jockey, 0)
	HorseSlice := make([]*model.Horse, 0)
	RHRSlice := make([]*model.RelationHorseRace, 0)

	resultTable.Find("tr").Each(func(i int, s *goquery.Selection) {
		//fmt.Printf("Content of cell %d: %s\n", i, s.Text())
		rhr := &model.RelationHorseRace{}
		horse := &model.Horse{}
		jockey := &model.Jockey{}
		s.Find("td").Each(func(j int, t *goquery.Selection) {
			//fmt.Printf("td contenc of %d: %s\n", j, t.Text())
			if j == 0 {
				rhr.Rank, err = strconv.ParseInt(t.Text(), 10, 64)
			}
			if j == 1 {
				rhr.FrameNumber, err = strconv.ParseInt(t.Text(), 10, 64)
			}
			if j == 2 {
				rhr.Number, err = strconv.ParseInt(t.Text(), 10, 64)
			}
			if j == 3 {
				href, ok := t.Children().Attr("href")
				if ok != true {
					fmt.Println("can't get horse href")
				}
				horseID, err := strconv.ParseInt(strings.Split(href, "/")[2], 10, 64)
				if err != nil {
					fmt.Println(err)
				}

				horse.Id = horseID
				horse.Name = t.Children().Text()
				horse.Link = href

				rhr.HorseId = horseID
			}
			if j == 4 {
				rhr.Sex = fmt.Sprint(t.Text()[0])
				rhr.Age, err = strconv.ParseInt(fmt.Sprint(t.Text()[1]), 10, 64)
			}
			if j == 5 {
				rhr.Handicap, err = strconv.ParseInt(t.Text(), 10, 64)
			}
			if j == 6 {
				href, ok := t.Children().Attr("href")
				if ok != true {
					fmt.Println(err)
				}
				jockeyID, err := strconv.ParseInt(strings.Split(href, "/")[2], 10, 64)
				if err != nil {
					fmt.Println(err)
				}

				jockey.Id = jockeyID
				jockey.Name = t.Children().Text()
				jockey.Link = href

				rhr.JockeyId = jockeyID
			}
			if j == 7 {
				rhr.GoalTime = t.Text()
			}
			if j == 11 {
				rhr.Final3F, err = strconv.ParseFloat(t.Text(), 64)
			}
			if j == 12 {
				rhr.Odds, err = strconv.ParseFloat(t.Text(), 64)
			}
			if j == 13 {
				rhr.Choice, err = strconv.ParseInt(t.Text(), 10, 64)
			}
			if j == 14 {
				rhr.HorseWeight, err = strconv.ParseInt(strings.Split(t.Text(), "(")[0], 10, 64)
				diff := strings.Split(strings.Split(t.Text(), "(")[1], ")")[0]
				rhr.WeightDiff, err = strconv.ParseInt(diff, 10, 64)
			}
		})
		JockeySlice = append(JockeySlice, jockey)
		HorseSlice = append(HorseSlice, horse)
		RHRSlice = append(RHRSlice, rhr)
	})

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		"conn-user",
		"password",
		"localhost:13306",
		"raceservice")
	db, err := sqlx.Open("mysql", dsn)
	jrepo := repository.NewJockeyRepository(db)
	hrepo := repository.NewHorseRepository(db)
	rhrrepo := repository.NewRhrRepository(db)

	ctx := context.Background()
	for _, jockey := range JockeySlice {
		if jockey.Id != 0 {
			status, err := jrepo.InsertJockey(ctx, jockey)
			if err != nil || status != 1 {
				fmt.Println(err)
			}
		}
	}
	for _, horse := range HorseSlice {
		if horse.Id != 0 {
			status, err := hrepo.InsertHorse(ctx, horse)
			if err != nil || status != 1 {
				fmt.Println(err)
			}
		}
	}
	for _, rhr := range RHRSlice {
		fmt.Println(rhr)
		if rhr.HorseId != 0 {
			status, err := rhrrepo.InsertRhr(ctx, rhr)
			if err != nil || status != 1 {
				fmt.Println(err)
			}
		}
	}

}
