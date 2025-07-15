package customErrors

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrGenerateSession = errors.New("failed to generate session id")
	ErrSetSession      = errors.New("failed to set session")
	ErrFindSession     = errors.New("failed to find session")
	ErrGetSession      = errors.New("failed to get session")
	ErrParseRedisValue = errors.New("failed to parse redis return value")
	ErrDeleteSession   = errors.New("failed to delete session")
	ErrMarshallData    = errors.New("failed to marshall data")
	ErrUnmarshallData  = errors.New("failed to unmarshall data")

	ErrDatabaseUser = errors.New("failed to make query to db")
	ErrNotUnique    = errors.New("data not unique")
	ErrCreateSalt   = errors.New("failed to generate salt")
)

func HandleAuthGRPCError(err error) error {
	if err == nil {
		return nil
	}

	st, ok := status.FromError(err)
	if !ok {
		return err
	}

	switch st.Code() {
	case codes.Unavailable:
		switch st.Message() {
		case "failed to generate session id":
			return ErrGenerateSession
		case "failed to set session":
			return ErrSetSession
		case "failed to get user id":
			return ErrGetSession
		default:
			return err
		}
	case codes.NotFound:
		switch st.Message() {
		case "failed to find user id by session id":
			return ErrFindSession
		case "failed to delete session":
			return ErrDeleteSession
		default:
			return err
		}
	case codes.FailedPrecondition:
		switch st.Message() {
		case "failed to marshall data":
			return ErrMarshallData
		case "failed to unmarshall data":
			return ErrUnmarshallData
		default:
			return err
		}
	case codes.InvalidArgument:
		return ErrParseRedisValue
	default:
		return err
	}
}

func HandleUserGRPCError(err error) error {
	if err == nil {
		return nil
	}

	st, ok := status.FromError(err)
	if !ok {
		return err
	}

	switch st.Code() {
	case codes.Unavailable:
		return ErrDatabaseUser
	case codes.InvalidArgument:
		return ErrNotUnique
	case codes.FailedPrecondition:
		return ErrCreateSalt
	default:
		return err
	}
}
