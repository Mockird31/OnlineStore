package domain

import (
    "context"
    "github.com/Mockird31/OnlineStore/internal/pkg/model"
)

type Usecase interface {
    SignupUser(ctx context.Context, regData *model.RegisterData) (*model.User, string, error)
}