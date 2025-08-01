package errorStatus

import (
	"net/http"

	"github.com/Mockird31/OnlineStore/internal/pkg/helpers/customErrors"
)

var mapErrorStatus = map[error]int{
	customErrors.ErrGenerateSession: http.StatusBadGateway,
	customErrors.ErrSetSession:      http.StatusServiceUnavailable,
	customErrors.ErrFindSession:     http.StatusServiceUnavailable,
	customErrors.ErrGetSession:      http.StatusBadRequest,
	customErrors.ErrParseRedisValue: http.StatusServiceUnavailable,
	customErrors.ErrDeleteSession:   http.StatusServiceUnavailable,
	customErrors.ErrMarshallData:    http.StatusServiceUnavailable,
	customErrors.ErrUnmarshallData:  http.StatusServiceUnavailable,

	customErrors.ErrDatabaseUser: http.StatusServiceUnavailable,
	customErrors.ErrNotUnique:    http.StatusBadRequest,
	customErrors.ErrCreateSalt:   http.StatusBadGateway,
}

func ErrorStatus(err error) int {
	status, exist := mapErrorStatus[err]
	if !exist {
		return http.StatusInternalServerError
	}
	return status
}
