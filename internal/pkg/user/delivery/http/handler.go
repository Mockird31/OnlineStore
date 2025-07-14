package http

import (
	"net/http"
	"time"

	"github.com/Mockird31/OnlineStore/internal/pkg/helpers/ctxWorker"
	"github.com/Mockird31/OnlineStore/internal/pkg/helpers/errorStatus"
	"github.com/Mockird31/OnlineStore/internal/pkg/helpers/json"
	loggerPkg "github.com/Mockird31/OnlineStore/internal/pkg/helpers/logger"
	"github.com/Mockird31/OnlineStore/internal/pkg/model"
	"github.com/Mockird31/OnlineStore/internal/pkg/user/domain"
	"github.com/go-playground/validator/v10"
)

const (
	SESSION_DURATION = 24 * time.Hour
)

type UserHandler struct {
	usecase domain.Usecase
}

func NewUserHandler(usecase domain.Usecase) *UserHandler {
	return &UserHandler{
		usecase: usecase,
	}
}

func (u *UserHandler) SignupUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := loggerPkg.LoggerFromContext(ctx)

	regData := &model.RegisterData{}

	err := json.ReadJSON(w, r, regData)
	if err != nil {
		logger.Error("failed to parse registration data")
		json.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse registration data", nil)
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	err = validate.Struct(regData)
	if err != nil {
		logger.Error("wrong register data")
		json.WriteErrorResponse(w, http.StatusBadRequest, "wrong register data", nil)
		return
	}

	user, sessionID, err := u.usecase.SignupUser(ctx, regData)
	if err != nil {
		logger.Error("failed to register user")
		json.WriteErrorResponse(w, errorStatus.ErrorStatus(err), "failed to register user", nil)
		return
	}

	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Expires:  time.Now().Add(SESSION_DURATION),
		Secure:   true,
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)
	json.WriteSuccessResponse(w, http.StatusOK, user, nil)
}

func (u *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := loggerPkg.LoggerFromContext(ctx)

	logData := &model.LoginData{}

	err := json.ReadJSON(w, r, logData)
	if err != nil {
		logger.Error("failed to parse registration data")
		json.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse registration data", nil)
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	err = validate.Struct(logData)
	if err != nil {
		logger.Error("wrong register data")
		json.WriteErrorResponse(w, http.StatusBadRequest, "wrong register data", nil)
		return
	}

	user, sessionID, err := u.usecase.LoginUser(ctx, logData)
	if err != nil {
		logger.Error("failed to register user")
		json.WriteErrorResponse(w, errorStatus.ErrorStatus(err), "failed to register user", nil)
		return
	}

	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Expires:  time.Now().Add(SESSION_DURATION),
		Secure:   true,
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)
	json.WriteSuccessResponse(w, http.StatusOK, user, nil)
}

func (u *UserHandler) LogoutUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := loggerPkg.LoggerFromContext(ctx)
	_, isAuth := ctxWorker.UserIDFromContext(ctx)
	if !isAuth {
		logger.Error("user already doesn't auth")
		json.WriteErrorResponse(w, http.StatusUnauthorized, "user not auth", nil)
		return
	}

	sessionCookie, err := r.Cookie("session_id")
	if err != nil {
		logger.Error("failed to get session cookie")
		json.WriteErrorResponse(w, http.StatusUnauthorized, "user not auth", nil)
		return
	}

	err = u.usecase.LogoutUser(ctx, sessionCookie.Value)
	if err != nil {
		logger.Error("failed to logout user")
		json.WriteErrorResponse(w, errorStatus.ErrorStatus(err), "failed to logout user", nil)
		return
	}

	expiredCookie := &http.Cookie{
		Name:     "session_id",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		Expires:  time.Now().Add(-24 * time.Hour),
		Secure:   true,
		HttpOnly: true,
	}

	http.SetCookie(w, expiredCookie)
	json.WriteSuccessResponse(w, http.StatusOK, "user was successfully logout", nil)
}
