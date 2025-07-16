package http

import (
	"net/http"
	"strconv"

	"github.com/Mockird31/OnlineStore/config"
	"github.com/Mockird31/OnlineStore/internal/pkg/helpers/json"
	loggerPkg "github.com/Mockird31/OnlineStore/internal/pkg/helpers/logger"
	"github.com/Mockird31/OnlineStore/internal/pkg/helpers/pagination"
	"github.com/Mockird31/OnlineStore/internal/pkg/item/domain"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type ItemHandler struct {
	itemUsecase domain.Usecase
	cfg         *config.Config
}

func NewItemHandler(usecase domain.Usecase, cfg *config.Config) *ItemHandler {
	return &ItemHandler{
		itemUsecase: usecase,
		cfg:         cfg,
	}
}

func (h *ItemHandler) GetItems(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := loggerPkg.LoggerFromContext(ctx)

	pagination, err := pagination.GetPagination(r, &h.cfg.PaginationConfig)
	if err != nil {
		logger.Error("failed to get pagination", zap.Error(err))
		json.WriteErrorResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	items, err := h.itemUsecase.GetItems(ctx, pagination)
	if err != nil {
		logger.Error("failed to get items", zap.Error(err))
		json.WriteErrorResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	json.WriteSuccessResponse(w, http.StatusOK, items, nil)
}

func (h *ItemHandler) GetItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := loggerPkg.LoggerFromContext(ctx)

	vars := mux.Vars(r)
	idString := vars["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		logger.Error("was entered wrong id item", zap.Error(err))
		json.WriteErrorResponse(w, http.StatusBadRequest, "wrong id", nil)
		return
	}

	item, err := h.itemUsecase.GetItem(ctx, int64(id))
	if err != nil {
		logger.Error("failed to get item", zap.Error(err))
		json.WriteErrorResponse(w, http.StatusServiceUnavailable, err.Error(), nil)
		return
	}

	json.WriteSuccessResponse(w, http.StatusOK, item, nil)
}
