package delivery

import (
	"context"

	authProto "github.com/Mockird31/OnlineStore/gen/auth"
	domain "github.com/Mockird31/OnlineStore/microservices/auth/internal/domain"
	models "github.com/Mockird31/OnlineStore/microservices/auth/models"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AuthService struct {
	authUsecase domain.Usecase
	authProto.UnimplementedAuthServiceServer
}

func NewAuthService(authUsecase domain.Usecase) authProto.AuthServiceServer {
	return &AuthService{authUsecase: authUsecase}
}

func (s *AuthService) CreateSession(ctx context.Context, userID *authProto.UserID) (*authProto.SessionID, error) {
	sessionID, err := s.authUsecase.CreateSession(ctx, models.UserIDToInt(userID))
	if err != nil {
		return nil, err
	}
	return models.StringToSessionID(sessionID), nil
}

func (s *AuthService) DeleteSession(ctx context.Context, sessionID *authProto.SessionID) (*emptypb.Empty, error) {
	err := s.authUsecase.DeleteSession(ctx, models.SessionIDToString(sessionID))
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *AuthService) GetUserIDBySessionID(ctx context.Context, sessionID *authProto.SessionID) (*authProto.UserID, error) {
	userID, err := s.authUsecase.GetUserIDBySessionID(ctx, models.SessionIDToString(sessionID))
	if err != nil {
		return nil, err
	}
	return models.IntToUserID(userID), nil
}
