package place

import (
	"context"
	"github.com/pkg/errors"
	"go01-airbnb/payload"
	"go01-airbnb/presenter"
	myerror "go01-airbnb/util/error"
	"gorm.io/gorm"
	"strings"
)

func (u *UseCase) Update(ctx context.Context,
	req *payload.UpdatePlaceRequest,
) (*presenter.PlaceResponseWrapper, error) {
	myPlace, err := u.PlaceRepo.GetByID(ctx, req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, myerror.ErrLocationNotFound()
		}

		return nil, myerror.ErrLocationGet(err)
	}

	if req.Name != nil {
		myPlace.Name = strings.TrimSpace(*req.Name)
	}

	err = u.PlaceRepo.Update(ctx, myPlace)
	if err != nil {
		return nil, myerror.ErrLocationUpdate(err)
	}

	return &presenter.PlaceResponseWrapper{
		Place: myPlace,
	}, nil
}
