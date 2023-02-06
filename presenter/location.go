package presenter

import "go01-airbnb/model"

type LocationResponseWrapper struct {
	Location *model.Location `json:"location"`
}

type ListLocationResponseWrapper struct {
	Location []model.Location `json:"locations"`
	Meta     interface{}      `json:"meta"`
}
