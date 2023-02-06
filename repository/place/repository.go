package place

import (
	"context"
	"go01-airbnb/codetype"
	"go01-airbnb/model"
)

type Repository interface {
	Create(ctx context.Context, data *model.Place) error
	GetByID(ctx context.Context, id int64) (*model.Place, error)
	GetList(ctx context.Context, paginator *codetype.Paginator) ([]model.Place, int64, error)
	Delete(ctx context.Context, data *model.Place) error
	Update(ctx context.Context, poll *model.Place) error
}
