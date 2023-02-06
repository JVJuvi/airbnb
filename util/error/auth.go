package myerror

import (
	"fmt"
	"net/http"
)

func ErrAuthGetUser(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "10200",
		Message:   "Failed to get User.",
	}
}

func ErrAuthCreateUser(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "10201",
		Message:   "Failed to create User.",
	}
}

func ErrAuthUpdateUser(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "10202",
		Message:   "Failed to update User.",
	}
}

func ErrAuthDeleteUser(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "10203",
		Message:   "Failed to delete User.",
	}
}

func ErrAuthNotFoundUser() MyError {
	return MyError{
		Raw:       nil,
		HTTPCode:  http.StatusNotFound,
		ErrorCode: "10204",
		Message:   "Not found.",
	}
}

func ErrAuthWrongCredential() MyError {
	return MyError{
		Raw:       nil,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: "10205",
		Message:   fmt.Sprintf("Invalid email or password"),
	}
}

func ErrAuthUnauthorized() MyError {
	return MyError{
		Raw:       nil,
		HTTPCode:  http.StatusUnauthorized,
		ErrorCode: "10206",
		Message:   fmt.Sprintf("Unauthorized!"),
	}
}

//func ErrValidate(param string) MyError {
//	return MyError{
//		Raw:       nil,
//		HTTPCode:  http.StatusBadRequest,
//		ErrorCode: "10207",
//		Message:   fmt.Sprintf("Fail to validate %s!", param),
//	}
//}

func ErrAuthInvalidRole() MyError {
	return MyError{
		Raw:       nil,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: "10208",
		Message:   fmt.Sprintf("Invalid role"),
	}
}

func ErrInvalidValueIsActive() MyError {
	return MyError{
		Raw:       nil,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: "10209",
		Message:   fmt.Sprintf("Invalid value isActive"),
	}
}

func ErrAuthUserExist() MyError {
	return MyError{
		Raw:       nil,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: "10210",
		Message:   "User is existed",
	}
}

func ErrAuthHashPassword(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "10211",
		Message:   "Can not hash password",
	}
}

func ErrAuthInvalidParams(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: "10212",
		Message:   "Invalid params.",
	}
}

func ErrAuthGenerateToken(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "10211",
		Message:   "Can not generate token",
	}
}

//func ErrStructValidation(param error) MyError {
//	for _, err := range param.(validator.ValidationErrors) {
//		return ErrValidate(err.Field())
//	}
//
//	return ErrInvalidValueIsActive()
//}
