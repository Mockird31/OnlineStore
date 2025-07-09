package domain

import (
	"context"

	"github.com/Mockird31/OnlineStore/microservices/user/model"
)

type Usecase interface {
	SignupUser(ctx context.Context, regData *model.RegisterData) (*model.User, error)
}
