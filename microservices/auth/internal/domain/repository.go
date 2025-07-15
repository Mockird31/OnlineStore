package domain

import (
	"context"
)

type Repository interface {
	CreateSession(ctx context.Context, user []byte) (string, error)
	GetUserBySessionID(ctx context.Context, sessionID string) ([]byte, error)
	DeleteSession(ctx context.Context, sessionID string) error
}
