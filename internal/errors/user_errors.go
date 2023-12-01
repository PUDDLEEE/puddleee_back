package errors

import "net/http"

var (
	INVALID_EMAIL    = NewError(http.StatusBadRequest, "invalid email", "101")
	EMAIL_EXISTED    = NewError(http.StatusBadRequest, "email already in use.", "102")
	USERNAME_EXISTED = NewError(http.StatusBadRequest, "username already in use.", "103")
)
