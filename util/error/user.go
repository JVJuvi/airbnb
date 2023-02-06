package myerror

import (
	"fmt"
	"net/http"
)

func ErrUserGet(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "10300", // required FE
		Message:   "Failed to get user.",
	}
}

func ErrUserUpdate(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "10302",
		Message:   "Failed to update user.",
	}
}

func ErrUserNotFound() MyError {
	return MyError{
		Raw:       nil,
		HTTPCode:  http.StatusNotFound,
		ErrorCode: "10304",
		Message:   "Not found user.",
	}
}

func ErrUserDuplicateParam(param string) MyError {
	return MyError{
		Raw:       nil,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: "10306",
		Message:   fmt.Sprintf("%s is existed", param),
	}
}

func ErrUserInvalidParam() MyError {
	return MyError{
		Raw:       nil,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: "10307",
		Message:   "Invalid param",
	}
}

func ErrUserInvalidRole() MyError {
	return MyError{
		Raw:       nil,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: "10307",
		Message:   "Invalid Role",
	}
}
