package payload

import (
	validation "go01-airbnb/client/validator"
	"go01-airbnb/codetype"
	myerror "go01-airbnb/util/error"
)

type RegisterRequest struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	FirstName string `json:"first_name" validate:"required,gte=0,lte=30"`
	LastName  string `json:"last_name" validate:"required,gte=0,lte=30"`
	Phone     string `json:"phone" validate:"required"`
	//Avatar    codetype.Image `json:"avatar"`
	Role     codetype.Role `json:"role" validate:"required"`
	IsActive bool          `json:"is_active" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (r *RegisterRequest) Validate() error {
	err := validation.GetValidator().Struct(r)
	if err != nil {
		return myerror.ErrStructValidation(err)
	}

	if !r.Role.IsValid() {
		return myerror.ErrAuthInvalidRole()
	}

	if !r.IsActive {
		return myerror.ErrInvalidValueIsActive()
	}

	return nil
}

func (l *LoginRequest) Validate() error {
	err := validation.GetValidator().Struct(l)
	if err != nil {
		return myerror.ErrStructValidation(err)
	}

	return nil
}
