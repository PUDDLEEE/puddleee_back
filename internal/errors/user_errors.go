package errors

import "net/http"

var (
	INVALID_EMAIL    = NewError(http.StatusBadRequest, "invalid email", "003")
	EMAIL_EXISTED    = NewError(http.StatusInternalServerError, "email already in use.", "004")
	USERNAME_EXISTED = NewError(http.StatusInternalServerError, "username already in use.", "005")
)
