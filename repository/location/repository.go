package location

import (
	"context"
	"go01-airbnb/codetype"
	"go01-airbnb/model"
)

type Repository interface {
	Create(ctx context.Context, data *model.Location) error
	GetByID(ctx context.Context, id int64) (*model.Location, error)
	GetList(ctx context.Context, paginator *codetype.Paginator) ([]model.Location, int64, error)
	Delete(ctx context.Context, data *model.Location) error
	Update(ctx context.Context, data *model.Location) error
}
