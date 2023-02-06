package place

import (
	"context"
	"go01-airbnb/model"
	"go01-airbnb/payload"
	"go01-airbnb/presenter"
	myerror "go01-airbnb/util/error"
)

func (u *UseCase) Create(
	ctx context.Context,
	req *payload.CreatePlaceRequest,
) (*presenter.PlaceResponseWrapper, error) {
	myPlace := &model.Place{
		Name:           req.Name,
		Description:    req.Description,
		TotalGuests:    req.TotalGuests,
		TotalBedrooms:  req.TotalBedrooms,
		TotalBathrooms: req.TotalBathrooms,
		PricePerNight:  req.PricePerNight,
		AverageRating:  req.AverageRating,
		OwnerID:        req.OwnerID,
		LocationID:     req.LocationID,
		Latitude:       req.Latitude,
		Longitude:      req.Longitude,
		Address:        req.Address,
	}

	err := u.PlaceRepo.Create(ctx, myPlace)
	if err != nil {
		return nil, myerror.ErrPlaceCreate(err)
	}

	return &presenter.PlaceResponseWrapper{
		Place: myPlace,
	}, nil
}
