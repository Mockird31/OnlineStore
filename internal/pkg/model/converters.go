package model

import (
	authProto "github.com/Mockird31/OnlineStore/gen/auth"
	userProto "github.com/Mockird31/OnlineStore/gen/user"
)

func IntToUserIDProto(userID int64) *authProto.UserID {
	return &authProto.UserID{Id: userID}
}

func SessionIDProtoToString(sessionID *authProto.SessionID) string {
	return sessionID.Id
}

func StringToSessionIDProto(sessionID string) *authProto.SessionID {
	return &authProto.SessionID{Id: sessionID}
}

func UserIDProtoToInt(userID *authProto.UserID) int64 {
	return userID.Id
}

func RegisterDataToProto(regData *RegisterData) *userProto.RegisterData {
	return &userProto.RegisterData{
		Username:        regData.Username,
		Email:           regData.Email,
		Password:        regData.Password,
		ConfirmPassword: regData.ConfirmPassword,
	}
}

func UserFromProto(user *userProto.User) *User {
	return &User{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
	}
}
