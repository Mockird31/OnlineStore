package auth

import (
	"context"
)

type Usecase interface {
	CreateSession(ctx context.Context, userID int64) (string, error)
	GetUserIDBySessionID(ctx context.Context, sessionID string) (int64, error)
	DeleteSession(ctx context.Context, sessionID string) error
}
