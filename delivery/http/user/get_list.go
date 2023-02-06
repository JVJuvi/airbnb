package user

import (
	"github.com/labstack/echo/v4"
	"go01-airbnb/payload"
	"go01-airbnb/presenter"
	"go01-airbnb/util/app"
	myerror "go01-airbnb/util/error"
)

func (r *Route) GetList(c echo.Context) error {
	var (
		ctx  = c.Request().Context()
		req  = payload.GetListUserRequest{}
		resp *presenter.ListUserResponseWrapper
	)

	if err := c.Bind(&req); err != nil {
		return app.Response.Error(c, myerror.ErrUserInvalidParam())
	}

	resp, err := r.UseCase.User.GetList(ctx, &req)
	if err != nil {
		return app.Response.Error(c, err.(myerror.MyError))
	}

	return app.Response.Success(c, resp)
}
