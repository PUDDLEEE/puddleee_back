package errors

import "net/http"

var (
	NOT_FOUND      = NewError(http.StatusNotFound, "resources not found", "001")
	INTERNAL_ERROR = NewError(http.StatusInternalServerError, "internal Server Error", "002")
	BAD_PARAM      = NewError(http.StatusBadRequest, "url param is incorrect.", "003")
	NOT_AUTHORIZED = NewError(http.StatusBadRequest, "not authorized.", "004")
)
