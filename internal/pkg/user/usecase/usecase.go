package usecase

import (
	"context"

	authProto "github.com/Mockird31/OnlineStore/gen/auth"
	userProto "github.com/Mockird31/OnlineStore/gen/user"
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
		return nil, "", err
	}
	user := model.UserFromProto(userProto)
	sessionIDProto, err := u.authClient.CreateSession(ctx, model.IntToUserIDProto(user.Id))
	if err != nil {
		return user, "", err
	}
	sessionID := model.SessionIDProtoToString(sessionIDProto)
	return user, sessionID, nil
}
