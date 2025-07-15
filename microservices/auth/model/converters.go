package model

import (
	authProto "github.com/Mockird31/OnlineStore/gen/auth"
)

func UserProtoToUser(user *authProto.User) *User {
	return &User{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
	}
}

func UserToUserProto(user *User) *authProto.User {
	return &authProto.User{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
	}
}

func StringToSessionID(sessionID string) *authProto.SessionID {
	return &authProto.SessionID{Id: sessionID}
}

func SessionIDToString(sessionID *authProto.SessionID) string {
	return sessionID.Id
}
