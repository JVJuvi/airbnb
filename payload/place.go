package payload

import (
	validation "go01-airbnb/client/validator"
	"go01-airbnb/codetype"
	myerror "go01-airbnb/util/error"
	"strings"
)

type CreatePlaceRequest struct {
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	TotalGuests    int     `json:"total_guests"`
	TotalBedrooms  int     `json:"total_bedrooms"`
	TotalBathrooms int     `json:"total_bathrooms"`
	PricePerNight  int     `json:"price_per_night"`
	AverageRating  float32 `json:"average_rating"`
	OwnerID        int     `json:"owner_id"`
	LocationID     int     `json:"location_id"`
	Latitude       float32 `json:"latitude"`
	Longitude      float32 `json:"longitude"`
	Address        string  `json:"address"`
}

type GetPlaceByIDRequest struct {
	ID int64 `json:"-"`
}

type GetListPlaceRequest struct {
	codetype.Paginator
}

func (g *GetListPlaceRequest) Format() {
	g.Paginator.Format()
}

type UpdatePlaceRequest struct {
	ID             int64   `json:"id"`
	Name           *string `json:"name"`
	Description    *string `json:"description"`
	TotalGuests    *int    `json:"total_guests"`
	TotalBedrooms  *int    `json:"total_bedrooms"`
	TotalBathrooms *int    `json:"total_bathrooms"`
	PricePerNight  *int    `json:"price_per_night"`
	AverageRating  float32 `json:"average_rating"`
	OwnerID        *int    `json:"owner_id"`
	LocationID     *int    `json:"location_id"`
	Latitude       float32 `json:"latitude"`
	Longitude      float32 `json:"longitude"`
	Address        string  `json:"address"`
}

type DeletePlaceRequest struct {
	ID int64 `json:"-"`
}

func (u *CreatePlaceRequest) Validate() error {
	u.Name = strings.TrimSpace(u.Name)
	u.Description = strings.TrimSpace(u.Description)
	u.Address = strings.TrimSpace(u.Address)
	err := validation.GetValidator().Struct(u)
	if err != nil {
		return myerror.ErrStructValidation(err)
	}

	return nil
}
