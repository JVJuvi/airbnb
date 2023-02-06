package place

import (
	"context"
	"github.com/pkg/errors"
	"go01-airbnb/payload"
	myerror "go01-airbnb/util/error"
	"gorm.io/gorm"
)

func (u *UseCase) Delete(ctx context.Context, req *payload.DeleteLocationRequest) error {
	myLocation, err := u.PlaceRepo.GetByID(ctx, req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return myerror.ErrLocationNotFound()
		}

		return myerror.ErrLocationGet(err)
	}

	err = u.PlaceRepo.Delete(ctx, myLocation)
	if err != nil {
		return myerror.ErrLocationDelete(err)
	}

	return nil
}
