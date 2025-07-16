package domain

import (
	"context"

	"github.com/Mockird31/OnlineStore/internal/pkg/model"
)

type Repository interface {
	GetItems(ctx context.Context, pagination *model.Pagination) ([]*model.Item, error)
	GetCategoriesByItemsID(ctx context.Context, itemsIDs []int64) (map[int64][]*model.Category, error)
	GetItem(ctx context.Context, itemID int64) (*model.Item, error)
	GetCategoriesByItemID(ctx context.Context, itemID int64) ([]*model.Category, error)
}
