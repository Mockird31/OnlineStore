package usecase

import (
	"context"

	authProto "github.com/Mockird31/OnlineStore/gen/auth"
	auth "github.com/Mockird31/OnlineStore/internal/pkg/auth/domain"
	loggerPkg "github.com/Mockird31/OnlineStore/internal/pkg/helpers/logger"
	"github.com/Mockird31/OnlineStore/internal/pkg/model"
)

type AuthUsecase struct {
	authClient authProto.AuthServiceClient
}

func NewAuthUsecase(authClient authProto.AuthServiceClient) auth.Usecase {
	return &AuthUsecase{authClient: authClient}
}

func (u *AuthUsecase) CreateSession(ctx context.Context, userID int64) (string, error) {
	logger := loggerPkg.LoggerFromContext(ctx)
	sessionIDProto, err := u.authClient.CreateSession(ctx, model.IntToUserIDProto(userID))
	if err != nil {
		logger.Error("failed to create session")
		return "", err
	}
	return model.SessionIDProtoToString(sessionIDProto), nil
}

func (u *AuthUsecase) GetUserIDBySessionID(ctx context.Context, sessionID string) (int64, error) {
	logger := loggerPkg.LoggerFromContext(ctx)
	userIDProto, err := u.authClient.GetUserIDBySessionID(ctx, model.StringToSessionIDProto(sessionID))
	if err != nil {
		logger.Error("failed to get session")
		return 0, nil
	}
	return model.UserIDProtoToInt(userIDProto), nil
}

func (u *AuthUsecase) DeleteSession(ctx context.Context, sessionID string) error {
	logger := loggerPkg.LoggerFromContext(ctx)
	_, err := u.authClient.DeleteSession(ctx, model.StringToSessionIDProto(sessionID))
	if err != nil {
		logger.Error("failed to delete session")
		return err
	}
	return nil
}
