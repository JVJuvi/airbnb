package auth

import (
	"context"
	"github.com/pkg/errors"
	"go01-airbnb/payload"
	"go01-airbnb/presenter"
	myerror "go01-airbnb/util/error"
	util "go01-airbnb/util/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (u *UseCase) Login(ctx context.Context,
	req *payload.LoginRequest,
) (*presenter.LoginResponseWrapper, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	user, err := u.UserRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, myerror.ErrAuthNotFoundUser()
		}

		return nil, myerror.ErrAuthGetUser(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, myerror.ErrAuthWrongCredential()
	}

	token, err := util.GenerateJWT(util.TokenPayload{Email: user.Email}, u.Config.App.Secret)
	if err != nil {
		return nil, myerror.ErrAuthGenerateToken(err)
	}

	return &presenter.LoginResponseWrapper{
		Jwt:  token.AccessToken,
		User: user,
	}, nil
}
