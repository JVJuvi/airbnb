package place

import (
	"context"
	"go01-airbnb/payload"
	"go01-airbnb/presenter"
)

type IUseCase interface {
	GetList(ctx context.Context, req *payload.GetListPlaceRequest) (*presenter.ListPlaceResponseWrapper, error)
	Create(ctx context.Context, req *payload.CreatePlaceRequest) (*presenter.PlaceResponseWrapper, error)
	Delete(ctx context.Context, req *payload.DeletePlaceRequest) error
	Update(ctx context.Context, req *payload.UpdatePlaceRequest) (*presenter.PlaceResponseWrapper, error)
	GetByID(ctx context.Context, req *payload.GetPlaceByIDRequest) (*presenter.PlaceResponseWrapper, error)
}
