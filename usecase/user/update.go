package user

import (
	"context"
	"github.com/pkg/errors"
	validation "go01-airbnb/client/validator"
	"go01-airbnb/payload"
	"go01-airbnb/presenter"
	myerror "go01-airbnb/util/error"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strings"
)

func (u *UseCase) Update(ctx context.Context, req *payload.UpdateUserRequest) (*presenter.UserResponseWrapper, error) {
	myUser, err := u.UserRepo.GetByID(ctx, req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, myerror.ErrUserNotFound()
		}

		return nil, myerror.ErrUserGet(err)
	}

	if req.Password != nil {
		*req.Password = strings.TrimSpace(*req.Password)

		err = bcrypt.CompareHashAndPassword([]byte(myUser.Password), []byte(*req.Password))
		if err == nil {
			return nil, myerror.ErrUserDuplicateParam("Password")
		}

		err = validation.GetValidator().Var(req.Password, "gte=5")
		if err != nil {
			return nil, myerror.ErrValidate(*req.Password)
		}

		myUser.Password = *req.Password
	}

	if req.Role != nil {
		if !req.Role.IsValid() {
			return nil, myerror.ErrUserInvalidRole()
		}

		myUser.Role = *req.Role
	}

	err = u.UserRepo.Update(ctx, myUser)
	if err != nil {
		return nil, myerror.ErrUserUpdate(err)
	}

	return &presenter.UserResponseWrapper{
		User: myUser,
	}, nil
}
