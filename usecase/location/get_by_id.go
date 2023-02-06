package location

import (
	"context"
	"github.com/pkg/errors"
	"go01-airbnb/payload"
	"go01-airbnb/presenter"
	myerror "go01-airbnb/util/error"
	"gorm.io/gorm"
)

func (u *UseCase) GetByID(
	ctx context.Context,
	req *payload.GetLocationByIDRequest,
) (*presenter.LocationResponseWrapper, error) {
	myLocation, err := u.LocationRepo.GetByID(ctx, req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, myerror.ErrLocationNotFound()
		}

		return nil, myerror.ErrLocationGet(err)
	}

	return &presenter.LocationResponseWrapper{
		Location: myLocation,
	}, nil
}
