package usecase

import (
	"context"
	"encoding/json"

	loggerPkg "github.com/Mockird31/OnlineStore/internal/pkg/helpers/logger"
	domain "github.com/Mockird31/OnlineStore/microservices/auth/internal/domain"
	"github.com/Mockird31/OnlineStore/microservices/auth/model"
	"github.com/Mockird31/OnlineStore/microservices/auth/model/errors"
)

type authUsecase struct {
	authRepository domain.Repository
}

func NewAuthUsecase(authRepository domain.Repository) domain.Usecase {
	return &authUsecase{
		authRepository: authRepository,
	}
}

func (u *authUsecase) CreateSession(ctx context.Context, user *model.User) (string, error) {
	logger := loggerPkg.LoggerFromContext(ctx)
	data, err := json.Marshal(user)
	if err != nil {
		logger.Error("failed to marshall user")
		return "", errors.NewMarshallDataError("failed to marshall data")
	}
	sessionID, err := u.authRepository.CreateSession(ctx, data)
	if err != nil {
		return "", err
	}
	return sessionID, nil
}

func (u *authUsecase) GetUserBySessionID(ctx context.Context, sessionID string) (*model.User, error) {
	logger := loggerPkg.LoggerFromContext(ctx)
	data, err := u.authRepository.GetUserBySessionID(ctx, sessionID)
	if err != nil {
		return nil, err
	}
	user := &model.User{}
	err = json.Unmarshal(data, user)
	if err != nil {
		logger.Error("failed to unmarshall user")
		return nil, errors.NewUnmarshallDataError("failed to unmarshall data")
	}
	return user, nil
}

func (u *authUsecase) DeleteSession(ctx context.Context, sessionID string) error {
	err := u.authRepository.DeleteSession(ctx, sessionID)
	if err != nil {
		return err
	}
	return nil
}
