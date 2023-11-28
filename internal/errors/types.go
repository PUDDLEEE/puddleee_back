package errors

import "net/http"

type CustomError struct {
	Code int
	Msg  string
	Data interface{}
}

var (
	INVALID_EMAIL  = NewError(http.StatusBadRequest, "Invalid email.")
	NOT_FOUND      = NewError(http.StatusNotFound, "Resources not found")
	INTERNAL_ERROR = NewError(http.StatusInternalServerError, "Internal Server Error")
)
