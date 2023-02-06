package place

import (
	"go01-airbnb/config"
	"go01-airbnb/repository"
	"go01-airbnb/repository/place"
)

type UseCase struct {
	PlaceRepo place.Repository
	Config    *config.Config
}

func New(repo *repository.Repository) IUseCase {
	return &UseCase{
		PlaceRepo: repo.Place,
		Config:    config.GetConfig(),
	}
}
