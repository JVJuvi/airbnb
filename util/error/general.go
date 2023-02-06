package myerror

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func ErrValidate(param string) MyError {
	return MyError{
		Raw:       nil,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: "001",
		Message:   fmt.Sprintf("Fail to validate %s!", param),
	}
}

func ErrNoPermission() MyError {
	return MyError{
		Raw:       nil,
		HTTPCode:  http.StatusForbidden,
		ErrorCode: "002",
		Message:   "No permission.",
	}
}

func ErrStructValidation(param error) MyError {
	for _, err := range param.(validator.ValidationErrors) {
		return ErrValidate(err.Field())
	}

	return ErrInvalidValueIsActive()
}
