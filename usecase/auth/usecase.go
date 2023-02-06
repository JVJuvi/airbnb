package auth

import (
	"go01-airbnb/config"
	"go01-airbnb/repository"
	"go01-airbnb/repository/user"
)

type UseCase struct {
	UserRepo user.Repository

	Config *config.Config
}

func New(repo *repository.Repository) IUseCase {
	return &UseCase{
		UserRepo: repo.User,
		Config:   config.GetConfig(),
	}
}
