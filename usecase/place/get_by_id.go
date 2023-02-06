package place

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
	req *payload.GetPlaceByIDRequest,
) (*presenter.PlaceResponseWrapper, error) {
	myPlace, err := u.PlaceRepo.GetByID(ctx, req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, myerror.ErrLocationNotFound()
		}

		return nil, myerror.ErrLocationGet(err)
	}

	return &presenter.PlaceResponseWrapper{
		Place: myPlace,
	}, nil
}
