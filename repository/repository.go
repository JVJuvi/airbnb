package repository

import (
	"context"
	"go01-airbnb/repository/location"
	"go01-airbnb/repository/place"
	"go01-airbnb/repository/user"
	"gorm.io/gorm"
)

type Repository struct {
	User     user.Repository
	Place    place.Repository
	Location location.Repository
}

func New(getClient func(ctx context.Context) *gorm.DB) *Repository {
	return &Repository{
		User:     user.NewPG(getClient),
		Place:    place.NewPG(getClient),
		Location: location.NewPG(getClient),
	}
}
