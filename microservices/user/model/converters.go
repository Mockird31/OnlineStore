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

func UserToUserProto(user *User) *userProto.User {
	userProto := &userProto.User{
		Username: user.Username,
		Email: user.Email,
		Id: user.Id,
	}
	return userProto
}