package errors

import "net/http"
	INVALID_EMAIL    = NewError(http.StatusBadRequest, "invalid email", "101")
	EMAIL_EXISTED    = NewError(http.StatusInternalServerError, "email already in use.", "102")
	USERNAME_EXISTED = NewError(http.StatusInternalServerError, "username already in use.", "103")
)
