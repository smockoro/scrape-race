package model

import "time"

type Race struct {
	Id         int64
	Distance   int64
	Racecource string
	Date       time.Time
}

type RelationHorseRace struct {
	RaceId      int64 `db:"race_id"`
	Rank        int64
	FrameNumber int64
	Number      int64
	HorseId     int64 `db:"horse_id"`
	Sex         string
	Age         int64
	Handicap    int64
	JockeyId    int64
	GoalTime    string
	Final3F     float64
	Odds        float64
	Choice      int64
	HorseWeight int64
	WeightDiff  int64
}

type Horse struct {
	Id   int64
	Name string
	Link string
}

type Jockey struct {
	Id   int64
	Name string
	Link string
}
