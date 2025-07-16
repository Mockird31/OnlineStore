package http

import (
	"net/http"
	"strconv"

	"github.com/Mockird31/OnlineStore/config"
	"github.com/Mockird31/OnlineStore/internal/pkg/category/domain"
	"github.com/Mockird31/OnlineStore/internal/pkg/helpers/json"
	loggerPkg "github.com/Mockird31/OnlineStore/internal/pkg/helpers/logger"
	"github.com/Mockird31/OnlineStore/internal/pkg/helpers/pagination"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type CategoryHandler struct {
	usecase domain.Usecase
	cfg     *config.Config
}

func NewCategoryHandler(usecase domain.Usecase, cfg *config.Config) *CategoryHandler {
	return &CategoryHandler{
		usecase: usecase,
		cfg:     cfg,
	}
}

func (h *CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := loggerPkg.LoggerFromContext(ctx)

	pagination, err := pagination.GetPagination(r, &h.cfg.PaginationConfig)
	if err != nil {
		logger.Error("failed to get pagination", zap.Error(err))
		json.WriteErrorResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	categories, err := h.usecase.GetCategories(ctx, pagination)
	if err != nil {
		logger.Error("failed to get categories", zap.Error(err))
		json.WriteErrorResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	json.WriteSuccessResponse(w, http.StatusOK, categories, nil)
}

func (h *CategoryHandler) GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := loggerPkg.LoggerFromContext(ctx)

	vars := mux.Vars(r)
	categoryIDString := vars["id"]
	categoryID, err := strconv.Atoi(categoryIDString)
	if err != nil {
		logger.Error("failed to get category id from vars", zap.Error(err))
		json.WriteErrorResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	category, err := h.usecase.GetCategoryByID(ctx, int64(categoryID))
	if err != nil {
		logger.Error("failed to get category by id", zap.Error(err))
		json.WriteErrorResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	json.WriteSuccessResponse(w, http.StatusOK, category, nil)
}
