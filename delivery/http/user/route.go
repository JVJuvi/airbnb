package user

import (
	"github.com/labstack/echo/v4"
	"go01-airbnb/usecase"
)

type Route struct {
	UseCase *usecase.UseCase
}

func Init(group *echo.Group, useCase *usecase.UseCase) {
	r := &Route{UseCase: useCase}

	group.GET("", r.GetList)
	group.GET("/:id", r.GetByID)
	group.PUT("/:id", r.Update)
	group.PUT("/:id/deactivate", r.UpdateInActiveUser)

}
