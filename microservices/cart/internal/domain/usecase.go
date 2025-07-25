package domain

import (
	"context"

	"github.com/Mockird31/OnlineStore/microservices/cart/model"
)

type Usecase interface {
	AddItem(ctx context.Context, req *model.AddRequest) error
	DeleteItem(ctx context.Context, req *model.DeleteRequest) error
	GetCart(ctx context.Context, req model.CartIdentifier) error
}
