package location

import (
	"context"
	"go01-airbnb/model"
	"go01-airbnb/payload"
	"go01-airbnb/presenter"
	myerror "go01-airbnb/util/error"
)

func (u *UseCase) Create(
	ctx context.Context,
	req *payload.CreateLocationRequest,
) (*presenter.LocationResponseWrapper, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	myLocation := &model.Location{
		Country: req.Country,
		State:   req.State,
		City:    req.City,
	}

	if err := u.LocationRepo.Create(ctx, myLocation); err != nil {
		return nil, myerror.ErrLocationCreate(err)
	}

	return &presenter.LocationResponseWrapper{
		Location: myLocation,
	}, nil
}
