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
	case codes.InvalidArgument:
		return ErrParseRedisValue
	default:
		return err
	}
}
