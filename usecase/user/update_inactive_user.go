package user

import (
	"context"
	"github.com/pkg/errors"
	"go01-airbnb/payload"
	"go01-airbnb/presenter"
	myerror "go01-airbnb/util/error"
	"gorm.io/gorm"
)

func (u *UseCase) UpdateInActiveUser(ctx context.Context, req *payload.UpdateUserStatusRequest) (*presenter.UserResponseWrapper, error) {
	myUser, err := u.UserRepo.GetByID(ctx, req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, myerror.ErrUserNotFound()
		}

		return nil, myerror.ErrUserGet(err)
	}

	if req.IsActive != nil {
		if myUser.IsActive == *req.IsActive {
			return &presenter.UserResponseWrapper{
				User: myUser,
			}, nil
		}

		myUser.IsActive = *req.IsActive
	}
	err = u.UserRepo.Update(ctx, myUser)
	if err != nil {
		return nil, myerror.ErrUserUpdate(err)
	}

	return &presenter.UserResponseWrapper{
		User: myUser,
	}, nil
}
