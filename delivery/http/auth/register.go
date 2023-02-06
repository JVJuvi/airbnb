package auth

import (
	"github.com/labstack/echo/v4"
	"go01-airbnb/payload"
	"go01-airbnb/presenter"
	"go01-airbnb/util/app"
	"go01-airbnb/util/customcontext"
	myerror "go01-airbnb/util/error"
)

func (r *Route) Register(c echo.Context) error {
	var (
		ctx  = &customcontext.CustomEchoContext{Context: c}
		req  = payload.RegisterRequest{}
		resp *presenter.UserResponseWrapper
	)

	if err := c.Bind(&req); err != nil {
		return app.Response.Error(c, myerror.ErrAuthInvalidParams(err))
	}

	err := r.UseCase.Auth.Register(ctx, &req)
	if err != nil {
		return app.Response.Error(c, err.(myerror.MyError))
	}

	return app.Response.Success(c, resp)
}
