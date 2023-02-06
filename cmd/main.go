package main

import (
	"github.com/labstack/gommon/log"
	"go01-airbnb/client/mysql"
	"go01-airbnb/delivery/http"
	"go01-airbnb/repository"
	"go01-airbnb/usecase"
)

func main() {
	client := mysql.GetClient
	repo := repository.New(client)
	useCase := usecase.New(repo)

	h := http.NewHTTPHandler(useCase, repo)
	if err := h.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
