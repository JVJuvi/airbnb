package user

import (
	"github.com/labstack/echo/v4"
	"go01-airbnb/payload"
	"go01-airbnb/presenter"
	"go01-airbnb/util/app"
	myerror "go01-airbnb/util/error"
	"strconv"
)

func (r *Route) Update(c echo.Context) error {
	var (
		ctx   = c.Request().Context()
		idStr = c.Param("id")
		req   = payload.UpdateUserRequest{}
		resp  *presenter.UserResponseWrapper
	)

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return app.Response.Error(c, myerror.ErrUserInvalidParam())
	}

	req = payload.UpdateUserRequest{ID: id}

	if err = c.Bind(&req); err != nil {
		return app.Response.Error(c, myerror.ErrUserInvalidParam())
	}

	resp, err = r.UseCase.User.Update(ctx, &req)
	if err != nil {
		return app.Response.Error(c, err.(myerror.MyError))
	}

	return app.Response.Success(c, resp)
}
