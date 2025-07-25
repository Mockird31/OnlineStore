package errors

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CartError struct {
	Code    codes.Code
	Message string
}

func (e *CartError) Error() string {
	return e.Message
}

func (e *CartError) GRPCStatus() *status.Status {
	return status.New(e.Code, e.Message)
}

func NewDatabaseError(format string, args ...interface{}) *CartError {
	return &CartError{
		Code:    codes.Unavailable,
		Message: fmt.Sprintf(format, args...),
	}
}

func NewCastError(format string, args ...interface{}) *CartError {
	return &CartError{
		Code:    codes.InvalidArgument,
		Message: fmt.Sprintf(format, args...),
	}
}
