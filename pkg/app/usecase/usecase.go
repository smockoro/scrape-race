package usecase

import (
	repo "github.com/smockoro/scrape-race/pkg/domain/repository"
)

type usecase struct {
	jrepo   repo.JockeyRepository
	hrepo   repo.HorseRepository
	rrepo   repo.RaceRepository
	rhrrepo repo.RhrRepository
}

func NewUsecase(jrepo repo.JockeyRepository, hrepo repo.HorseRepository,
	rrepo repo.RaceRepository, rhrrepo repo.RhrRepository) *usecase {
	return &usecase{
		jrepo:   jrepo,
		hrepo:   hrepo,
		rrepo:   rrepo,
		rhrrepo: rhrrepo,
	}
}
