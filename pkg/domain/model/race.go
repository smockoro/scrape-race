package model

import "time"

type Race struct {
	Id         int64
	Distance   int64
	Racecource string
	Date       time.Time
}

type RelationHorseRace struct {
	RaceId      int64
	Rank        int64
	FrameNumber int64
	Number      int64
	HorseId     int64
	Sex         string
	Age         int64
	Handicap    int64
	JockeyId    int64
	GoalTime    time.Time
	Final3F     time.Time
	Odds        float64
	HorseWeight int64
}

type Horse struct {
	Id        int64
	HorseName string
	HorseLink string
}

type Jockey struct {
	Id   int64
	Name string
}
