package errors

type CustomError struct {
	ErrorCode  string      `json:"error_code"`
	StatusCode int         `json:"status_code"`
	Msg        string      `json:"msg"`
	Data       interface{} `json:"data,omitempty"`
}
