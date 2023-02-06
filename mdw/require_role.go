package mdw

import (
	"github.com/labstack/echo/v4"
	"go01-airbnb/model"
	"go01-airbnb/util/app"
	myerror "go01-airbnb/util/error"
)

func RequiredRole() func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user, ok := c.Get(UserKey).(*model.User)
			if !ok {
				return app.Response.Error(c, myerror.ErrAuthUnauthorized())
			}

			if user.Role.IsAdmin() {
				return next(c)
			}

			return app.Response.Error(c, myerror.ErrNoPermission())
		}
	}
}
