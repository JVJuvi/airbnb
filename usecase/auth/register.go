package auth

import (
	"context"
	"github.com/pkg/errors"
	"go01-airbnb/model"
	"go01-airbnb/payload"
	myerror "go01-airbnb/util/error"
	util "go01-airbnb/util/jwt"
	"gorm.io/gorm"
)

func (u *UseCase) Register(ctx context.Context, req *payload.RegisterRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}

	user, err := u.UserRepo.GetByEmail(ctx, req.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return myerror.ErrAuthGetUser(err)
	}

	if user != nil {
		return myerror.ErrAuthUserExist()
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return myerror.ErrAuthHashPassword(err)
	}

	myUser := model.User{
		Email:     req.Email,
		Password:  hashedPassword,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Phone:     req.Phone,
		//Avatar:    req.Avatar,
		Role:     req.Role,
		IsActive: req.IsActive,
	}

	if err = u.UserRepo.Create(ctx, &myUser); err != nil {
		return myerror.ErrAuthGetUser(err)
	}

	return nil
}
