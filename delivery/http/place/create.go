package place

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
		req  = payload.CreatePlaceRequest{}
		resp *presenter.PlaceResponseWrapper
	)

	if err := c.Bind(&req); err != nil {
		return app.Response.Error(c, myerror.ErrLocationInvalidParam())
	}

	resp, err := r.UseCase.Place.Create(ctx, &req)
	if err != nil {
		return app.Response.Error(c, err.(myerror.MyError))
	}

	return app.Response.Success(c, resp)
}
