package location

import (
	"context"
	"github.com/pkg/errors"
	"go01-airbnb/model"
	"go01-airbnb/payload"
	"go01-airbnb/presenter"
	myerror "go01-airbnb/util/error"
	"gorm.io/gorm"
	"strings"
)

func (u *UseCase) Update(ctx context.Context,
	req *payload.UpdateLocationRequest,
) (*presenter.LocationResponseWrapper, error) {
	myLocation, err := u.LocationRepo.GetByID(ctx, req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, myerror.ErrLocationNotFound()
		}

		return nil, myerror.ErrLocationGet(err)
	}

	location, err := Validate(req, myLocation)
	if err != nil {
		return nil, err
	}

	err = u.LocationRepo.Update(ctx, location)
	if err != nil {
		return nil, myerror.ErrLocationUpdate(err)
	}

	return &presenter.LocationResponseWrapper{
		Location: location,
	}, nil
}

func Validate(req *payload.UpdateLocationRequest, myLocation *model.Location) (*model.Location, error) {
	if req.Country != nil {
		if *req.Country == myLocation.Country {
			return nil, myerror.ErrLocationDuplicateValue("country")
		}

		myLocation.Country = strings.TrimSpace(*req.Country)
	}

	if req.State != nil {
		if *req.State == myLocation.State {
			return nil, myerror.ErrLocationDuplicateValue("state")
		}

		myLocation.State = strings.TrimSpace(*req.State)
	}

	if req.City != nil {
		if *req.City == myLocation.City {
			return nil, myerror.ErrLocationDuplicateValue("city")
		}

		myLocation.City = strings.TrimSpace(*req.City)
	}

	return myLocation, nil
}
