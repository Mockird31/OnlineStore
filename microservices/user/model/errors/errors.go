package errors

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserError struct {
	Code    codes.Code
	Message string
}

func (e *UserError) Error() string {
	return e.Message
}

func (e *UserError) GRPCStatus() *status.Status {
	return status.New(e.Code, e.Message)
}

func NewDatabaseError(format string, args ...interface{}) *UserError {
	return &UserError{
		Code:    codes.Unavailable,
		Message: fmt.Sprintf(format, args...),
	}
}

func NewNotUniqueError(format string, args ...interface{}) *UserError {
	return &UserError{
		Code:    codes.InvalidArgument,
		Message: fmt.Sprintf(format, args...),
	}
}

func NewCreateSaltError(format string, args ...interface{}) *UserError {
	return &UserError{
		Code:    codes.FailedPrecondition,
		Message: fmt.Sprintf(format, args...),
	}
}

func NewUserNotExistError(format string, args ...interface{}) *UserError {
	return &UserError{
		Code:    codes.InvalidArgument,
		Message: fmt.Sprintf(format, args...),
	}
}

func NewWrongPasswordError(format string, args ...interface{}) *UserError {
	return &UserError{
		Code:    codes.InvalidArgument,
		Message: fmt.Sprintf(format, args...),
	}
}
