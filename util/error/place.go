package myerror

import (
	"fmt"
	"net/http"
)

func ErrPlaceGet(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "10400", // required FE
		Message:   "Failed to get place.",
	}
}

func ErrPlaceUpdate(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "10402",
		Message:   "Failed to update place.",
	}
}

func ErrPlaceNotFound() MyError {
	return MyError{
		Raw:       nil,
		HTTPCode:  http.StatusNotFound,
		ErrorCode: "10404",
		Message:   "Not found place.",
	}
}

func ErrPlaceDuplicateParam(param string) MyError {
	return MyError{
		Raw:       nil,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: "10406",
		Message:   fmt.Sprintf("%s is existed", param),
	}
}

func ErrPlaceInvalidParam() MyError {
	return MyError{
		Raw:       nil,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: "10407",
		Message:   "Invalid param",
	}
}

func ErrPlaceUnauthorized() MyError {
	return MyError{
		Raw:       nil,
		HTTPCode:  http.StatusUnauthorized,
		ErrorCode: "10408",
		Message:   "Unauthorized user",
	}
}

func ErrPlaceCreate(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "10409",
		Message:   "Failed to create place.",
	}
}
