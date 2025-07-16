package domain

import (
	"context"

	"github.com/Mockird31/OnlineStore/internal/pkg/model"
)

type Repository interface {
	GetCategories(ctx context.Context, filters *model.Pagination) ([]*model.Category, error)
	GetCategoryByID(ctx context.Context, id int64) (*model.Category, error)
}
