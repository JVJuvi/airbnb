package location

import (
	"context"
	"go01-airbnb/payload"
	"go01-airbnb/presenter"
)

type IUseCase interface {
	GetList(ctx context.Context, req *payload.GetListLocationRequest) (*presenter.ListLocationResponseWrapper, error)
	Create(ctx context.Context, req *payload.CreateLocationRequest) (*presenter.LocationResponseWrapper, error)
	Delete(ctx context.Context, req *payload.DeleteLocationRequest) error
	Update(ctx context.Context, req *payload.UpdateLocationRequest) (*presenter.LocationResponseWrapper, error)
	GetByID(ctx context.Context, req *payload.GetLocationByIDRequest) (*presenter.LocationResponseWrapper, error)
}
