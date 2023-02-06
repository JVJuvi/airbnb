package user

import (
	"context"
	"go01-airbnb/codetype"
	"go01-airbnb/model"
)

type Repository interface {
	GetByInterface(ctx context.Context, conditions interface{}) (*model.User, error)
	GetByID(ctx context.Context, id int64) (*model.User, error)
	GetList(ctx context.Context, paginator *codetype.Paginator) ([]model.User, int64, error)
	Update(ctx context.Context, user *model.User) error
	Create(ctx context.Context, data *model.User) error
	GetByEmail(ctx context.Context, email string) (*model.User, error)
}
