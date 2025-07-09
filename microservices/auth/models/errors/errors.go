package errors

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthError struct {
	Code    codes.Code
	Message string
}

func (e *AuthError) Error() string {
	return e.Message
}

func (e *AuthError) GRPCStatus() *status.Status {
	return status.New(e.Code, e.Message)
}

func NewGenerateSessionError(format string, args ...interface{}) *AuthError {
	return &AuthError{
		Code:    codes.Internal,
		Message: fmt.Sprintf(format, args...),
	}
}

func NewSetSessionIDError(format string, args ...interface{}) *AuthError {
	return &AuthError{
		Code:    codes.Unavailable,
		Message: fmt.Sprintf(format, args...),
	}
}

func NewFindSessionError(format string, args ...interface{}) *AuthError {
	return &AuthError{
		Code:    codes.NotFound,
		Message: fmt.Sprintf(format, args...),
	}
}

func NewGetSessionError(format string, args ...interface{}) *AuthError {
	return &AuthError{
		Code:    codes.Unavailable,
		Message: fmt.Sprintf(format, args...),
	}
}

func NewFailToParseRedisIntError(format string, args ...interface{}) *AuthError {
	return &AuthError{
		Code:    codes.InvalidArgument,
		Message: fmt.Sprintf(format, args...),
	}
}

func NewDeleteSessionError(format string, args ...interface{}) *AuthError {
	return &AuthError{
		Code:    codes.NotFound,
		Message: fmt.Sprintf(format, args...),
	}
}
