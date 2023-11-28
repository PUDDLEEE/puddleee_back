package errors

import "net/http"

var (
	NOT_FOUND      = NewError(http.StatusNotFound, "resources not found", "001")
	INTERNAL_ERROR = NewError(http.StatusInternalServerError, "internal Server Error", "002")
)
