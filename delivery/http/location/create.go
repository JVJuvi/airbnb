package location

import (
	"github.com/labstack/echo/v4"
	"go01-airbnb/payload"
	"go01-airbnb/presenter"
	"go01-airbnb/util/app"
	myerror "go01-airbnb/util/error"
)

func (r *Route) Create(c echo.Context) error {
	var (
		ctx  = c.Request().Context()
		req  = payload.CreateLocationRequest{}
		resp *presenter.LocationResponseWrapper
	)

	if err := c.Bind(&req); err != nil {
		return app.Response.Error(c, myerror.ErrLocationInvalidParam())
	}

	resp, err := r.UseCase.Location.Create(ctx, &req)
	if err != nil {
		return app.Response.Error(c, err.(myerror.MyError))
	}

	return app.Response.Success(c, resp)
}
