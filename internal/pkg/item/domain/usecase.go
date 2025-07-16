package domain

import (
	"context"

	"github.com/Mockird31/OnlineStore/internal/pkg/model"
)

type Usecase interface {
	GetItems(ctx context.Context, pagination *model.Pagination) ([]*model.Item, error)
	GetItem(ctx context.Context, itemID int64) (*model.Item, error)
}
