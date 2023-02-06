package model

import "time"

type Location struct {
	ID        int       `json:"id"`
	Country   string    `json:"country"`
	State     string    `json:"state"`
	City      string    `json:"city"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
