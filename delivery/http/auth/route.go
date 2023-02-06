package auth

import (
	"github.com/labstack/echo/v4"
	"go01-airbnb/usecase"
)

type Route struct {
	UseCase *usecase.UseCase
}

func Init(group *echo.Group, useCase *usecase.UseCase) {
	r := &Route{UseCase: useCase}

	group.POST("/login", r.Login)
	group.POST("/register", r.Register)
}
