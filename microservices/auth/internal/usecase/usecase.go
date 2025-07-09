package usecase

import (
	"context"

	domain "github.com/Mockird31/OnlineStore/microservices/auth/internal/domain"
)

type authUsecase struct {
	authRepository domain.Repository
}

func NewAuthUsecase(authRepository domain.Repository) domain.Usecase {
	return &authUsecase{
		authRepository: authRepository,
	}
}

func (u *authUsecase) CreateSession(ctx context.Context, userID int64) (string, error) {
	sessionID, err := u.authRepository.CreateSession(ctx, userID)
	if err != nil {
		return "", err
	}
	return sessionID, nil
}

func (u *authUsecase) GetUserIDBySessionID(ctx context.Context, sessionID string) (int64, error) {
	userID, err := u.authRepository.GetUserIDBySessionID(ctx, sessionID)
	if err != nil {
		return userID, err
	}
	return userID, nil
}

func (u *authUsecase) DeleteSession(ctx context.Context, sessionID string) error {
	err := u.authRepository.DeleteSession(ctx, sessionID)
	if err != nil {
		return err
	}
	return nil
}
