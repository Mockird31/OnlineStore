package model

import (
	userProto "github.com/Mockird31/OnlineStore/gen/user"
)

func RegisterDataProtoToRegisterData(regDataProto *userProto.RegisterData) *RegisterData {
	regData := &RegisterData{
		Username:        regDataProto.Username,
		Email:           regDataProto.Email,
		Password:        regDataProto.Password,
		ConfirmPassword: regDataProto.ConfirmPassword,
	}
	return regData
}

func LoginDataProtoToLoginData(loginDataProto *userProto.LoginData) *LoginData {
	logData := &LoginData{
		Username: loginDataProto.Username,
		Email:    loginDataProto.Email,
		Password: loginDataProto.Password,
	}
	return logData
}

func UserToUserProto(user *User) *userProto.User {
	userProto := &userProto.User{
		Username: user.Username,
		Email:    user.Email,
		Id:       user.Id,
	}
	return userProto
}
