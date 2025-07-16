package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Mockird31/OnlineStore/internal/pkg/category/domain"
	loggerPkg "github.com/Mockird31/OnlineStore/internal/pkg/helpers/logger"
	"github.com/Mockird31/OnlineStore/internal/pkg/model"
	"go.uber.org/zap"
)

const (
	GetCategoriesQuery = `
		SELECT id, title
		FROM category
		LIMIT $1 OFFSET $2
	`
	GetCategoryByIDQuery = `
		SELECT title
		FROM category
		WHERE id = $1
	`
)

type categoryPostgresRepository struct {
	db *sql.DB
}

func NewCategoryPostgresRepository(db *sql.DB) domain.Repository {
	return &categoryPostgresRepository{
		db: db,
	}
}

func (r *categoryPostgresRepository) GetCategories(ctx context.Context, filters *model.Pagination) ([]*model.Category, error) {
	logger := loggerPkg.LoggerFromContext(ctx)
	fmt.Println(filters)
	rows, err := r.db.QueryContext(ctx, GetCategoriesQuery, filters.Limit, filters.Offset)
	if err != nil {
		logger.Error("failed to get categories")
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			logger.Error("failed to close rows", zap.Error(err))
		}
	}()

	categories := make([]*model.Category, 0)
	for rows.Next() {
		var category model.Category
		err := rows.Scan(&category.Id, &category.Title)
		if err != nil {
			logger.Error("failed to scan category")
			return nil, err
		}
		categories = append(categories, &category)
	}

	if err := rows.Err(); err != nil {
		logger.Error("failed to get categories", zap.Error(err))
		return nil, err
	}

	return categories, nil
}

func (r *categoryPostgresRepository) GetCategoryByID(ctx context.Context, id int64) (*model.Category, error) {
	logger := loggerPkg.LoggerFromContext(ctx)
	var category model.Category
	category.Id = id
	err := r.db.QueryRowContext(ctx, GetCategoryByIDQuery, id).Scan(&category.Title)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			logger.Error("no categories with this id")
			return nil, sql.ErrNoRows
		}
		logger.Error("failed to get category by id")
		return nil, err
	}
	return &category, nil
}
