package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/smockoro/scrape-race/pkg/domain/model"
	repo "github.com/smockoro/scrape-race/pkg/domain/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type jockeyRepository struct {
	db *sqlx.DB
}
type horseRepository struct {
	db *sqlx.DB
}
type raceRepository struct {
	db *sqlx.DB
}
type rhrRepository struct {
	db *sqlx.DB
}

// NewJockeyRepository :
func NewJockeyRepository(db *sqlx.DB) repo.JockeyRepository {
	return &jockeyRepository{
		db: db,
	}
}

// NewHorseRepository :
func NewHorseRepository(db *sqlx.DB) repo.HorseRepository {
	return &horseRepository{
		db: db,
	}
}

// NewRaceRepository :
func NewRaceRepository(db *sqlx.DB) repo.RaceRepository {
	return &raceRepository{
		db: db,
	}
}

// NewRhrRepository :
func NewRhrRepository(db *sqlx.DB) repo.RhrRepository {
	return &rhrRepository{
		db: db,
	}
}

func (u *jockeyRepository) InsertJockey(ctx context.Context, jockey *model.Jockey) (int64, error) {
	res, err := u.db.NamedExecContext(ctx,
		"INSERT INTO jockeys(`id`, `name`, `link_url`) VALUES(:id, :name, :link)",
		jockey)
	if err != nil {
		return -1, status.Error(codes.Unknown, "failed to insert jockey"+err.Error())
	}

	id, err := res.LastInsertId()
	if err != nil {
		return -1, status.Error(codes.Unknown, "failed to retrieve jockey id"+err.Error())
	}

	return id, nil
}

func (u *horseRepository) InsertHorse(ctx context.Context, horse *model.Horse) (int64, error) {
	res, err := u.db.NamedExecContext(ctx,
		"INSERT INTO horses(`id`, `name`, `link_url`) VALUES(:id, :name, :link)",
		horse)
	if err != nil {
		return -1, status.Error(codes.Unknown, "failed to insert horse"+err.Error())
	}

	id, err := res.LastInsertId()
	if err != nil {
		return -1, status.Error(codes.Unknown, "failed to retrieve horse id"+err.Error())
	}

	return id, nil
}

func (u *raceRepository) InsertRace(ctx context.Context, race *model.Race) (int64, error) {
	res, err := u.db.NamedExecContext(ctx,
		"INSERT INTO races(`name`, `age`, `mail`, `address`) VALUES(:name, :age, :mail, :address)",
		race)
	if err != nil {
		return -1, status.Error(codes.Unknown, "failed to insert race"+err.Error())
	}

	id, err := res.LastInsertId()
	if err != nil {
		return -1, status.Error(codes.Unknown, "failed to retrieve race id"+err.Error())
	}

	return id, nil
}

func (u *rhrRepository) InsertRhr(ctx context.Context, rhr *model.RelationHorseRace) (int64, error) {
	res, err := u.db.NamedExecContext(ctx,
		"INSERT INTO relation_horse_race(`race_id`, `horse_id`) VALUES(:race_id, :horse_id)",
		rhr)
	if err != nil {
		return -1, status.Error(codes.Unknown, "failed to insert rhr"+err.Error())
	}

	id, err := res.LastInsertId()
	if err != nil {
		return -1, status.Error(codes.Unknown, "failed to retrieve rhr id"+err.Error())
	}

	return id, nil
}

func (u *rhrRepository) SelectHorseID(ctx context.Context, id int64) ([]*model.RelationHorseRace, error) {
	return nil, nil
}

func (u *rhrRepository) SelectRaceID(ctx context.Context, id int64) ([]*model.RelationHorseRace, error) {
	return nil, nil
}
