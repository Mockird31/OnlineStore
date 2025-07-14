package domain

import (
	"context"

	"github.com/Mockird31/OnlineStore/microservices/user/model"
)

type Repository interface {
	SignupUser(ctx context.Context, username, email, passwordHash string) (*model.User, error)
	CheckUsernameUnique(ctx context.Context, username string) (bool, error) 
	CheckEmailUnique(ctx context.Context, email string) (bool, error)
	GetPasswordHash(ctx context.Context, username, email string) (string, error)
	GetUserIDByUsername(ctx context.Context, username string) (int64, error)
}
