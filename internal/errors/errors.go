package errors

func (e *CustomError) Error() string {
	return e.Msg

}
func NewError(code int, msg string) *CustomError {
	return &CustomError{
		Msg:  msg,
		Code: code,
	}
}

func GetError(e *CustomError, data interface{}) *CustomError {
	return &CustomError{
		Msg:  e.Msg,
		Code: e.Code,
		Data: e.Data,
	}
}
