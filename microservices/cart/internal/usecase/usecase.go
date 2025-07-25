package usecase

import (
	"context"

	loggerPkg "github.com/Mockird31/OnlineStore/internal/pkg/helpers/logger"
	"github.com/Mockird31/OnlineStore/microservices/cart/internal/domain"
	"github.com/Mockird31/OnlineStore/microservices/cart/model"
	"github.com/Mockird31/OnlineStore/microservices/cart/model/errors"
)

type IdentifierType int

const (
	StateInt IdentifierType = iota
	StateString
)

type cartUsecase struct {
	cartPostgresRepository domain.Repository
}

func NewCartUsecase(cartPostgresRepository domain.Repository) domain.Usecase {
	return &cartUsecase{cartPostgresRepository: cartPostgresRepository}
}

func GetCartIdentifier(ctx context.Context, cartIdentifier model.CartIdentifier) (interface{}, IdentifierType, error) {
	logger := loggerPkg.LoggerFromContext(ctx)
	if cartIdentifier.IsCartIdentifier() > 0 {
		userIDStruct, ok := cartIdentifier.(*model.UserID)
		if !ok {
			logger.Error("failed to cast")
			return nil, -1, errors.NewCastError("failed to cast to cart id")
		}
		return userIDStruct.UserID, StateInt
	} else {
		sessionIDStruct, ok := cartIdentifier.(*model.SessionID)
		if !ok {
			logger.Error("failed to cast")
			return nil, -1, errors.NewCastError("failed to cast to cart id")
		}
		return sessionIDStruct.SessionID, StateString, nil
	}
}

func (u *cartUsecase) GetCartID(ctx context.Context, cartIdentifier model.CartIdentifier) (int64, error) {
	logger := loggerPkg.LoggerFromContext(ctx)
	var ok bool
	var cartID int64
	var cartIDByUserInt int64
	var cartIDByUserString string
	var err error
	cartIDByUser, idType, err := GetCartIdentifier(ctx, cartIdentifier)
	if err != nil {
		return 0, err
	}

	if idType == StateInt {
		cartIDByUserInt, ok = cartIDByUser.(int64)
		if !ok {
			logger.Error("failed to cast interface to int")
			return 0, errors.NewCastError("failed to cast")
		}
	} else {
		cartIDByUserString, ok = cartIDByUser.(string)
		if !ok {
			logger.Error("failed to cast interface to string")
			return 0, errors.NewCastError("failed to cast")
		}
	}

	if idType == StateInt {
		cartID, err = u.cartPostgresRepository.GetCartIDByUserID(ctx, cartIDByUserInt)
		if err != nil {
			return 0, err
		}
	} else {
		cartID, err = u.cartPostgresRepository.GetCartIDBySessionID(ctx, cartIDByUserString)
		if err != nil {
			return 0, err
		}
	}
	return cartID, nil
}

func (u *cartUsecase) AddItem(ctx context.Context, req *model.AddRequest) error {
	cartID, err := u.GetCartID(ctx, req.CartID)
	if err != nil {
		return err
	}

	err = u.cartPostgresRepository.AddItem(ctx, cartID, req.ItemID)
	if err != nil {
		return err
	}
	return nil
}
