package usecase

import (
	"context"

	"github.com/Mockird31/OnlineStore/internal/pkg/category/domain"
	"github.com/Mockird31/OnlineStore/internal/pkg/model"
)

type CategoryUsecase struct {
	categoryRepository domain.Repository
}

func NewCategoryUsecase(categoryRepository domain.Repository) domain.Usecase {
	return &CategoryUsecase{
		categoryRepository: categoryRepository,
	}
}

func (u *CategoryUsecase) GetCategories(ctx context.Context, filters *model.Pagination) ([]*model.Category, error) {
	categories, err := u.categoryRepository.GetCategories(ctx, filters)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (u *CategoryUsecase) GetCategoryByID(ctx context.Context, id int64) (*model.Category, error) {
	category, err := u.categoryRepository.GetCategoryByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return category, nil
}
