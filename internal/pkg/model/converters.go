package model

import (
	authProto "github.com/Mockird31/OnlineStore/gen/auth"
	userProto "github.com/Mockird31/OnlineStore/gen/user"
)

func UserToAuthUser(user *User) *authProto.User {
	return &authProto.User{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
	}
}

func AuthProtoUserToUser(user *authProto.User) *User {
	return &User{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
	}
}

func SessionIDProtoToString(sessionID *authProto.SessionID) string {
	return sessionID.Id
}

func StringToSessionIDProto(sessionID string) *authProto.SessionID {
	return &authProto.SessionID{Id: sessionID}
}

func RegisterDataToProto(regData *RegisterData) *userProto.RegisterData {
	return &userProto.RegisterData{
		Username:        regData.Username,
		Email:           regData.Email,
		Password:        regData.Password,
		ConfirmPassword: regData.ConfirmPassword,
	}
}

func LoginDataToProto(logData *LoginData) *userProto.LoginData {
	return &userProto.LoginData{
		Username: logData.Username,
		Email:    logData.Email,
		Password: logData.Password,
	}
}

func UserFromProto(user *userProto.User) *User {
	return &User{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
	}
}
