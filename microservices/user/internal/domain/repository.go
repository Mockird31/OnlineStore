package domain

import (
	"context"

	"github.com/Mockird31/OnlineStore/microservices/user/model"
)

type Repository interface {
	SignupUser(ctx context.Context, regData *model.RegisterData) *model.User
}
