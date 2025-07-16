package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	loggerPkg "github.com/Mockird31/OnlineStore/internal/pkg/helpers/logger"
	"github.com/Mockird31/OnlineStore/internal/pkg/item/domain"
	"github.com/Mockird31/OnlineStore/internal/pkg/model"
	"go.uber.org/zap"
)

const (
	GetItemsQuery = `
		SELECT id, title, description, price, image_url, count, is_active
		FROM item
		LIMIT $1 OFFSET $2
	`
	GetItemQuery = `
		SELECT title, description, price, image_url, count, is_active
		FROM item 
		WHERE id = $1
	`
	GetCategoriesByItemIDQuery = `
		SELECT c.id, c.title
		FROM item_category ic
		JOIN category c ON ic.category_id = c.id
		WHERE ic.item_id = $1
	`
)

type itemPostgresRepository struct {
	db *sql.DB
}

func NewItemPostgresRepository(db *sql.DB) domain.Repository {
	return &itemPostgresRepository{
		db: db,
	}
}

func (r *itemPostgresRepository) GetItems(ctx context.Context, pagination *model.Pagination) ([]*model.Item, error) {
	logger := loggerPkg.LoggerFromContext(ctx)

	rows, err := r.db.QueryContext(ctx, GetItemsQuery, pagination.Limit, pagination.Offset)
	if err != nil {
		logger.Error("failed to create query to db", zap.Error(err))
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			logger.Error("failed to close rows")
		}
	}()

	items := make([]*model.Item, 0)
	for rows.Next() {
		var item model.Item
		err := rows.Scan(&item.Id, &item.Title, &item.Description, &item.Price, &item.ImageURL, &item.Count, &item.IsActive)
		if err != nil {
			logger.Error("failed to scan row", zap.Error(err))
			return nil, err
		}
		items = append(items, &item)
	}

	if err := rows.Err(); err != nil {
		logger.Error("failed to get items", zap.Error(err))
		return nil, err
	}
	return items, nil
}

func (r *itemPostgresRepository) GetCategoriesByItemsID(ctx context.Context, itemsIDs []int64) (map[int64][]*model.Category, error) {
	logger := loggerPkg.LoggerFromContext(ctx)

	if len(itemsIDs) == 0 {
		return make(map[int64][]*model.Category), nil
	}

	placeholders := make([]string, len(itemsIDs))
	args := make([]interface{}, len(itemsIDs))
	for i, id := range itemsIDs {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
		args[i] = id
	}

	query := fmt.Sprintf(`
        SELECT ic.item_id, c.id, c.title
        FROM item_category ic
        JOIN category c ON ic.category_id = c.id
        WHERE ic.item_id IN (%s)
        ORDER BY ic.item_id, c.id
    `, strings.Join(placeholders, ", "))

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		logger.Error("failed to get categories for items", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	categoriesByItemID := make(map[int64][]*model.Category)

	for _, itemID := range itemsIDs {
		categoriesByItemID[itemID] = make([]*model.Category, 0)
	}

	for rows.Next() {
		var itemID int64
		var category model.Category

		if err := rows.Scan(&itemID, &category.Id, &category.Title); err != nil {
			logger.Error("failed to scan category", zap.Error(err))
			return nil, err
		}

		categoriesByItemID[itemID] = append(categoriesByItemID[itemID], &category)
	}

	if err := rows.Err(); err != nil {
		logger.Error("error iterating rows", zap.Error(err))
		return nil, err
	}

	return categoriesByItemID, nil
}

func (r *itemPostgresRepository) GetItem(ctx context.Context, itemID int64) (*model.Item, error) {
	logger := loggerPkg.LoggerFromContext(ctx)
	var item model.Item
	err := r.db.QueryRowContext(ctx, GetItemQuery, itemID).Scan(&item.Title, &item.Description, &item.Price, &item.ImageURL, &item.Count, &item.IsActive)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			logger.Error("this item not exist")
			return nil, sql.ErrNoRows
		}
		logger.Error("failed to get item")
		return nil, err
	}

	item.Id = itemID
	return &item, nil
}

func (r *itemPostgresRepository) GetCategoriesByItemID(ctx context.Context, itemID int64) ([]*model.Category, error) {
	logger := loggerPkg.LoggerFromContext(ctx)
	rows, err := r.db.QueryContext(ctx, GetCategoriesByItemIDQuery, itemID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			logger.Error("categories not exist for this item")
			return nil, sql.ErrNoRows
		}
		logger.Error("failed to get categories")
		return nil, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			logger.Error("failed to close rows")
		}
	}()

	categories := make([]*model.Category, 0)
	for rows.Next() {
		var category model.Category
		err := rows.Scan(&category.Id, &category.Title)
		if err != nil {
			logger.Error("failed to scan row")
			return nil, err
		}
		categories = append(categories, &category)
	}

	if err := rows.Err(); err != nil {
		logger.Error("error iterating rows", zap.Error(err))
		return nil, err
	}

	return categories, nil
}
