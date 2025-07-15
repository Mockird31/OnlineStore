package delivery

import (
	"context"

	authProto "github.com/Mockird31/OnlineStore/gen/auth"
	domain "github.com/Mockird31/OnlineStore/microservices/auth/internal/domain"
	model "github.com/Mockird31/OnlineStore/microservices/auth/model"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AuthService struct {
	authUsecase domain.Usecase
	authProto.UnimplementedAuthServiceServer
}

func NewAuthService(authUsecase domain.Usecase) authProto.AuthServiceServer {
	return &AuthService{authUsecase: authUsecase}
}

func (s *AuthService) CreateSession(ctx context.Context, user *authProto.User) (*authProto.SessionID, error) {
	sessionID, err := s.authUsecase.CreateSession(ctx, model.UserProtoToUser(user))
	if err != nil {
		return nil, err
	}
	return model.StringToSessionID(sessionID), nil
}

func (s *AuthService) DeleteSession(ctx context.Context, sessionID *authProto.SessionID) (*emptypb.Empty, error) {
	err := s.authUsecase.DeleteSession(ctx, model.SessionIDToString(sessionID))
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *AuthService) GetUserBySessionID(ctx context.Context, sessionID *authProto.SessionID) (*authProto.User, error) {
	user, err := s.authUsecase.GetUserBySessionID(ctx, model.SessionIDToString(sessionID))
	if err != nil {
		return nil, err
	}
	return model.UserToUserProto(user), nil
}
