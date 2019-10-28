package repository

import (
	"context"

	"github.com/smockoro/scrape-race/pkg/domain/model"
)

type JockeyRepository interface {
	InsertJockey(context.Context, *model.Jockey) (int64, error)
}

type HorseRepository interface {
	InsertHorse(context.Context, *model.Horse) (int64, error)
}

type RaceRepository interface {
	InsertRace(context.Context, *model.Race) (int64, error)
}

// RhrRepository : Rhr -> Relation of Horse and Race
type RhrRepository interface {
	InsertRhr(context.Context, *model.RelationHorseRace) (int64, error)
	SelectRaceID(context.Context, int64) ([]*model.RelationHorseRace, error)
	SelectHorseID(context.Context, int64) ([]*model.RelationHorseRace, error)
}
