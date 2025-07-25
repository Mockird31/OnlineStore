package repository

import (
	"context"

	"database/sql"

	loggerPkg "github.com/Mockird31/OnlineStore/internal/pkg/helpers/logger"
	"github.com/Mockird31/OnlineStore/microservices/cart/internal/domain"
	customErrors "github.com/Mockird31/OnlineStore/microservices/cart/model/errors"
)

const (
	AddItemQuery = `
	INSERT INTO cart_item (cart_id, item_id)
	VALUES ($1, $2)
	ON CONFLICT (cart_id, item_id)
	DO UPDATE SET quantity = cart_item.quantity + 1
	`
	GetCartIDByUserIDQuery = `
	SELECT id 
	FROM cart
	WHERE user_id = $1
	`
	GetCartIDBySessionIDQuery = `
	SELECT id 
	FROM cart
	WHERE session_id = $1
	`
)

type cartPostgresRepository struct {
	db *sql.DB
}

func NewCartPostgresRepository(db *sql.DB) domain.Repository {
	return &cartPostgresRepository{
		db: db,
	}
}

func (r *cartPostgresRepository) GetCartIDByUserID(ctx context.Context, userID int64) (int64, error) {
	logger := loggerPkg.LoggerFromContext(ctx)

	var cartID int64

	err := r.db.QueryRowContext(ctx, GetCartIDByUserIDQuery, userID).Scan(&cartID)
	if err != nil {
		logger.Error("failed to get cart")
		return 0, customErrors.NewDatabaseError("failed to do database query")
	}

	return cartID, nil
}

func (r *cartPostgresRepository) GetCartIDBySessionID(ctx context.Context, sessionID string) (int64, error) {
	logger := loggerPkg.LoggerFromContext(ctx)

	var cartID int64

	err := r.db.QueryRowContext(ctx, GetCartIDBySessionIDQuery, sessionID).Scan(&cartID)
	if err != nil {
		logger.Error("failed to get cart")
		return 0, customErrors.NewDatabaseError("failed to do database query")
	}

	return cartID, nil
}

func (r *cartPostgresRepository) AddItem(ctx context.Context, cartID int64, itemID int64) error {
	logger := loggerPkg.LoggerFromContext(ctx)

	_, err := r.db.ExecContext(ctx, AddItemQuery, cartID, itemID)
	if err != nil {
		logger.Error("failed to add item in cart")
		return customErrors.NewDatabaseError("failed to do database query")
	}

	return nil
}
