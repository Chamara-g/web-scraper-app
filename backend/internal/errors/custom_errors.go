package errors

import "fmt"

type CustomError struct {
	Code    int
	Message string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func New(code int, message string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}