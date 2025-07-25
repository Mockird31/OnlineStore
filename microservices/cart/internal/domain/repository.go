package domain

import (
	"context"
)

type Repository interface {
	GetCartIDByUserID(ctx context.Context, userID int64) (int64, error)
	GetCartIDBySessionID(ctx context.Context, sessionID string) (int64, error)
	AddItem(ctx context.Context, cartID int64, itemID int64) error
}
