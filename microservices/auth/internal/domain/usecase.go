package domain

import (
	"context"

	"github.com/Mockird31/OnlineStore/microservices/auth/model"
)

type Usecase interface {
	CreateSession(ctx context.Context, user *model.User) (string, error)
	GetUserBySessionID(ctx context.Context, sessionID string) (*model.User, error)
	DeleteSession(ctx context.Context, sessionID string) error
}
