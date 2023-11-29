package errors

// CustomError is the model for custom error detail
type CustomError struct {
	// ErrorCode holds more information that error occurred.
	ErrorCode string `json:"error_code"`

	// StatusCode holds http response status.
	StatusCode int `json:"status_code"`

	// Msg holds error message.
	Msg string `json:"msg"`

	// Data holds error data.
	Data interface{} `json:"data,omitempty"`
}
