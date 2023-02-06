package model

import (
	"go01-airbnb/codetype"
	"time"
)

type User struct {
	ID        int           `json:"id"`
	Email     string        `json:"email"`
	Password  string        `json:"-"`
	FirstName string        `json:"first_name"`
	LastName  string        `json:"last_name"`
	Phone     string        `json:"phone"`
	Role      codetype.Role `json:"role"`
	IsActive  bool          `json:"is_active"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	DeleteAt  time.Time     `json:"delete_at"`
}
