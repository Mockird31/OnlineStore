package delivery

import (
	"context"

	cartProto "github.com/Mockird31/OnlineStore/gen/cart"
	"github.com/Mockird31/OnlineStore/microservices/cart/model"
)

type UserService struct {
	cartProto.UnimplementedCartServiceServer
	userUsecase domain.Usecase
}

