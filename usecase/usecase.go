package usecase

import (
	"go01-airbnb/repository"
	"go01-airbnb/usecase/auth"
	"go01-airbnb/usecase/location"
	"go01-airbnb/usecase/place"
	"go01-airbnb/usecase/user"
)

type UseCase struct {
	User     user.IUseCase
	Auth     auth.IUseCase
	Location location.IUseCase
	Place    place.IUseCase
}

func New(repo *repository.Repository) *UseCase {
	return &UseCase{
		User:     user.New(repo),
		Auth:     auth.New(repo),
		Location: location.New(repo),
		Place:    place.New(repo),
	}
}
