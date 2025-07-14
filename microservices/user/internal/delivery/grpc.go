package delivery

import (
	"context"

	userProto "github.com/Mockird31/OnlineStore/gen/user"
	"github.com/Mockird31/OnlineStore/microservices/user/internal/domain"
	"github.com/Mockird31/OnlineStore/microservices/user/model"
)

type UserService struct {
	userProto.UnimplementedUserServiceServer
	userUsecase domain.Usecase
}

func NewUserService(userUsecase domain.Usecase) *UserService {
	return &UserService{userUsecase: userUsecase}
}

func (s *UserService) SignupUser(ctx context.Context, regData *userProto.RegisterData) (*userProto.User, error) {
	user, err := s.userUsecase.SignupUser(ctx, model.RegisterDataProtoToRegisterData(regData))
	if err != nil {
		return nil, err
	}

	return model.UserToUserProto(user), nil
}

func (s *UserService) LoginUser(ctx context.Context, logData *userProto.LoginData) (*userProto.User, error) {
	user, err := s.userUsecase.LoginUser(ctx, model.LoginDataProtoToLoginData(logData))
	if err != nil {
		return nil, err
	}
	return model.UserToUserProto(user), nil
}
