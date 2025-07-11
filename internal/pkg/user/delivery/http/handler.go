package http

import (
	"net/http"

	"github.com/Mockird31/OnlineStore/internal/pkg/helpers/json"
	loggerPkg "github.com/Mockird31/OnlineStore/internal/pkg/helpers/logger"
	"github.com/Mockird31/OnlineStore/internal/pkg/model"
	"github.com/Mockird31/OnlineStore/internal/pkg/user/domain"
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
}
