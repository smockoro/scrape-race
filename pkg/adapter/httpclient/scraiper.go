package main

import (
	"bufio"
	"fmt"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
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
		s.Find("td").Each(func(j int, t *goquery.Selection) {
			//fmt.Printf("td contenc of %d: %s\n", j, t.Text())
			if j == 0 {
				fmt.Printf("着順 : %s", t.Text())
				rhr.Rank, err = strconv.ParseInt(t.Text(), 10, 64)
			}
			if j == 1 {
				fmt.Printf(" 枠番 : %s", t.Text())
				rhr.FrameNumber, err = strconv.ParseInt(t.Text(), 10, 64)
			}
			if j == 2 {
				fmt.Printf(" 馬番 : %s", t.Text())
				rhr.Number, err = strconv.ParseInt(t.Text(), 10, 64)
			}
			if j == 3 {
				fmt.Printf(" 馬名 : %s", t.Children().Text())
				href, ok := t.Children().Attr("href")
				if ok != true {
					fmt.Println(err)
				}
				fmt.Printf(" 馬名リンク : %s", href)
			}
			if j == 4 {
				fmt.Printf(" 性齢 : %s", t.Text())
				rhr.Sex = fmt.Sprint(t.Text()[0])
				rhr.Age, err = strconv.ParseInt(fmt.Sprint(t.Text()[1]), 10, 64)
			}
			if j == 5 {
				fmt.Printf(" 斤量 : %s", t.Text())
				rhr.Handicap, err = strconv.ParseInt(t.Text(), 10, 64)
			}
			if j == 6 {
				fmt.Printf(" 騎手 : %s", t.Children().Text())
				href, ok := t.Children().Attr("href")
				if ok != true {
					fmt.Println(err)
				}
				fmt.Printf(" 騎手リンク : %s", href)
			}
			if j == 7 {
				fmt.Printf(" タイム : %s", t.Text())
			}
			if j == 11 {
				fmt.Printf(" 上り : %s", t.Text())
			}
			if j == 13 {
				fmt.Printf(" 人気 : %s", t.Text())
			}
			if j == 14 {
				fmt.Printf(" 馬体重 : %s\n", t.Text())
			}
		})
	})
}
