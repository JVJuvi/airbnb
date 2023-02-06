package model

import "time"

type Place struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	TotalGuests    int       `json:"total_guests"`
	TotalBedrooms  int       `json:"total_bedrooms"`
	TotalBathrooms int       `json:"total_bathrooms"`
	PricePerNight  int       `json:"price_per_night"`
	AverageRating  float32   `json:"average_rating"`
	OwnerID        int       `json:"owner_id"`
	LocationID     int       `json:"location_id"`
	Latitude       float32   `json:"latitude"`
	Longitude      float32   `json:"longitude"`
	Address        string    `json:"address"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
