package presenter

import "go01-airbnb/model"

type UserResponseWrapper struct {
	User *model.User `json:"user"`
}

type ListUserResponseWrapper struct {
	User []model.User `json:"users"`
	Meta interface{}  `json:"meta"`
}
