package usecase

import (
	"context"

	loggerPkg "github.com/Mockird31/OnlineStore/internal/pkg/helpers/logger"
	"github.com/Mockird31/OnlineStore/internal/pkg/item/domain"
	"github.com/Mockird31/OnlineStore/internal/pkg/model"
	"go.uber.org/zap"
)

type ItemUsecase struct {
	itemRepository domain.Repository
}

func NewItemUsecase(itemRepository domain.Repository) domain.Usecase {
	return &ItemUsecase{
		itemRepository: itemRepository,
	}
}

func (u *ItemUsecase) GetItems(ctx context.Context, pagination *model.Pagination) ([]*model.Item, error) {
	logger := loggerPkg.LoggerFromContext(ctx)

	items, err := u.itemRepository.GetItems(ctx, pagination)
	if err != nil {
		logger.Error("failed to get items", zap.Error(err))
		return nil, err
	}

	if len(items) == 0 {
		return items, nil
	}

	itemsIDs := make([]int64, len(items))
	for i, item := range items {
		itemsIDs[i] = item.Id
	}

	categoriesByItemID, err := u.itemRepository.GetCategoriesByItemsID(ctx, itemsIDs)
	if err != nil {
		logger.Error("failed to get categories for items", zap.Error(err))
		return nil, err
	}

	for i, item := range items {
		if categories, ok := categoriesByItemID[item.Id]; ok {
			items[i].Categories = categories
		} else {
			items[i].Categories = make([]*model.Category, 0)
		}
	}

	return items, nil
}

func (u *ItemUsecase) GetItem(ctx context.Context, itemID int64) (*model.Item, error) {
	item, err := u.itemRepository.GetItem(ctx, itemID)
	if err != nil {
		return nil, err
	}

	categories, err := u.itemRepository.GetCategoriesByItemID(ctx, itemID)
	if err != nil {
		return nil, err
	}
	
	item.Categories = categories
	return item, nil
}