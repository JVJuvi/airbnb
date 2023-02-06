package http

import (
	"github.com/labstack/echo/v4"
	"go01-airbnb/delivery/http/auth"
	"go01-airbnb/delivery/http/location"
	"go01-airbnb/delivery/http/place"
	"go01-airbnb/delivery/http/user"
	"go01-airbnb/mdw"
	"go01-airbnb/repository"
	"go01-airbnb/usecase"
)

func NewHTTPHandler(useCase *usecase.UseCase, repo *repository.Repository) *echo.Echo {
	e := echo.New()

	// APIs
	auth.Init(e.Group(""), useCase)

	api := e.Group("/api", mdw.RequiredAuth(repo))

	user.Init(api.Group("/users", mdw.RequiredRole()), useCase)
	location.Init(api.Group("/locations"), useCase)
	place.Init(api.Group("/places"), useCase)

	return e
}
