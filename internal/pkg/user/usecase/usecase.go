package usecase

import (
	"context"

	authProto "github.com/Mockird31/OnlineStore/gen/auth"
	userProto "github.com/Mockird31/OnlineStore/gen/user"
	"github.com/Mockird31/OnlineStore/internal/pkg/helpers/customErrors"
	"github.com/Mockird31/OnlineStore/internal/pkg/model"
	userDomain "github.com/Mockird31/OnlineStore/internal/pkg/user/domain"
)

type userUsecase struct {
	userClient userProto.UserServiceClient
	authClient authProto.AuthServiceClient
}

func NewUserUsecase(userClient userProto.UserServiceClient, authClient authProto.AuthServiceClient) userDomain.Usecase {
	return &userUsecase{userClient: userClient,
		authClient: authClient}
}

func (u *userUsecase) SignupUser(ctx context.Context, regData *model.RegisterData) (*model.User, string, error) {
	userProto, err := u.userClient.SignupUser(ctx, model.RegisterDataToProto(regData))
	if err != nil {
		return nil, "", customErrors.HandleUserGRPCError(err)
	}
	user := model.UserFromProto(userProto)
	sessionIDProto, err := u.authClient.CreateSession(ctx, model.IntToUserIDProto(user.Id))
	if err != nil {
		return user, "", customErrors.HandleAuthGRPCError(err)
	}
	sessionID := model.SessionIDProtoToString(sessionIDProto)
	return user, sessionID, nil
}

func (u *userUsecase) LoginUser(ctx context.Context, logData *model.LoginData) (*model.User, string, error) {
	userProto, err := u.userClient.LoginUser(ctx, model.LoginDataToProto(logData))
	if err != nil {
		return nil, "", customErrors.HandleUserGRPCError(err)
	}
	user := model.UserFromProto(userProto)
	sessionIDProto, err := u.authClient.CreateSession(ctx, model.IntToUserIDProto(user.Id))
	if err != nil {
		return user, "", customErrors.HandleAuthGRPCError(err)
	}
	sessionID := model.SessionIDProtoToString(sessionIDProto)
	return user, sessionID, nil
}

func (u *userUsecase) LogoutUser(ctx context.Context, sessionID string) error {
	_, err := u.authClient.DeleteSession(ctx, model.StringToSessionIDProto(sessionID))
	if err != nil {
		return customErrors.HandleAuthGRPCError(err)
	}
	return nil
}