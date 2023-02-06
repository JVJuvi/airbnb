package presenter

import "go01-airbnb/model"

type LoginResponseWrapper struct {
	Jwt  string      `json:"jwt"`
	User *model.User `json:"user"`
}
