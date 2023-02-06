package mdw

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"go01-airbnb/config"
	"go01-airbnb/repository"
	"go01-airbnb/util/app"
	"go01-airbnb/util/customcontext"
	myerror "go01-airbnb/util/error"
	util "go01-airbnb/util/jwt"
	"net/http"
	"strings"
)

const UserKey = "USER"

func extractTokenFromHeader(r *http.Request) (string, error) {
	bearerToken := r.Header.Get(echo.HeaderAuthorization)

	temp := strings.Split(bearerToken, " ")

	if temp[0] != "Bearer" || len(temp) < 2 || strings.TrimSpace(temp[1]) == "" {
		return "", errors.New("invalid token")
	}

	return temp[1], nil
}

func RequiredAuth(repo *repository.Repository) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cfg := config.GetConfig()
			ctx := customcontext.CustomEchoContext{Context: c}
			token, err := extractTokenFromHeader(c.Request())
			if err != nil {
				return app.Response.Error(c, myerror.ErrAuthUnauthorized())
			}

			payload, err := util.ValidateJWT(token, cfg)
			if err != nil {
				return app.Response.Error(c, myerror.ErrAuthUnauthorized())
			}

			//check user exist
			user, err := repo.User.GetByEmail(c.Request().Context(), payload.Email)
			if err != nil {
				return app.Response.Error(c, myerror.ErrAuthUnauthorized())
			}

			if !user.IsActive {
				return app.Response.Error(c, myerror.ErrAuthUnauthorized())
			}

			ctx.Set(UserKey, user)

			return next(c)
		}
	}
}
