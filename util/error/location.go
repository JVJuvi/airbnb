package myerror

import (
	"fmt"
	"net/http"
)

func ErrLocationGet(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "10500", // required FE
		Message:   "Failed to get location.",
	}
}

func ErrLocationUpdate(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "10502",
		Message:   "Failed to update location.",
	}
}

func ErrLocationNotFound() MyError {
	return MyError{
		Raw:       nil,
		HTTPCode:  http.StatusNotFound,
		ErrorCode: "10504",
		Message:   "Not found location.",
	}
}

func ErrLocationDuplicateValue(param string) MyError {
	return MyError{
		Raw:       nil,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: "10506",
		Message:   fmt.Sprintf("Duplicate %s", param),
	}
}

func ErrLocationInvalidParam() MyError {
	return MyError{
		Raw:       nil,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: "10507",
		Message:   "Invalid param",
	}
}

func ErrLocationUnauthorized() MyError {
	return MyError{
		Raw:       nil,
		HTTPCode:  http.StatusUnauthorized,
		ErrorCode: "10508",
		Message:   "Unauthorized user",
	}
}

func ErrLocationCreate(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "10509",
		Message:   "Failed to create location.",
	}
}

func ErrLocationDelete(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "10510",
		Message:   "Failed to delete location.",
	}
}
