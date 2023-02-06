package user

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
	req *payload.GetUserByIDRequest,
) (*presenter.UserResponseWrapper, error) {
	myUser, err := u.UserRepo.GetByID(ctx, req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, myerror.ErrUserNotFound()
		}

		return nil, myerror.ErrUserGet(err)
	}

	return &presenter.UserResponseWrapper{
		User: myUser,
	}, nil
}
