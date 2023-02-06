package presenter

import "go01-airbnb/model"

type PlaceResponseWrapper struct {
	Place *model.Place `json:"place"`
}

type ListPlaceResponseWrapper struct {
	Place []model.Place `json:"places"`
	Meta  interface{}   `json:"meta"`
}
