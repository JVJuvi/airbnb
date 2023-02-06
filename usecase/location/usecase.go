package location

import (
	"go01-airbnb/config"
	"go01-airbnb/repository"
	"go01-airbnb/repository/location"
)

type UseCase struct {
	LocationRepo location.Repository
	Config       *config.Config
}

func New(repo *repository.Repository) IUseCase {
	return &UseCase{
		LocationRepo: repo.Location,
		Config:       config.GetConfig(),
	}
}
