package payload

import (
	validation "go01-airbnb/client/validator"
	"go01-airbnb/codetype"
	myerror "go01-airbnb/util/error"
	"net/mail"
)

type GetUserByIDRequest struct {
	ID int64 `json:"-"`
}

type GetListUserRequest struct {
	codetype.Paginator
}

func (g *GetListUserRequest) Format() {
	g.Paginator.Format()
}

type UpdateUserRequest struct {
	ID        int64          `json:"-"`
	Email     string         `json:"email" validate:"required,email"`
	Password  *string        `json:"password" validate:"required"`
	FirstName string         `json:"first_name" validate:"required,gte=0,lte=30"`
	LastName  string         `json:"last_name" validate:"required,gte=0,lte=30"`
	Phone     string         `json:"phonße" validate:"required"`
	Role      *codetype.Role `json:"role" validate:"required"`
	IsActive  bool           `json:"is_active" validate:"required"`
}

type UpdateUserStatusRequest struct {
	ID       int64 `json:"-"`
	IsActive *bool `json:"is_active"`
}

type DeleteUserRequest struct {
	ID int64 `json:"-"`
}

func (u *UpdateUserRequest) Validate() error {
	err := validation.GetValidator().Struct(u)
	if err != nil {
		return myerror.ErrStructValidation(err)
	}

	return nil
}

func (u *UpdateUserRequest) ValidMailAddress(address string) bool {
	_, err := mail.ParseAddress(address)
	if err != nil {
		return false
	}
	return true
}
