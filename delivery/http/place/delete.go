package place

import (
	"github.com/labstack/echo/v4"
	"go01-airbnb/payload"
	"go01-airbnb/util/app"
	myerror "go01-airbnb/util/error"
	"strconv"
)

func (r *Route) Delete(c echo.Context) error {
	var (
		ctx   = c.Request().Context()
		idStr = c.Param("id")
		req   = payload.DeletePlaceRequest{}
	)

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return app.Response.Error(c, myerror.ErrLocationInvalidParam())
	}

	req.ID = id

	err = r.UseCase.Place.Delete(ctx, &req)
	if err != nil {
		return app.Response.Error(c, err.(myerror.MyError))
	}

	return app.Response.Success(c, nil)
}
