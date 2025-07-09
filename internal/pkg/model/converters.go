package model

import (
	authProto "github.com/Mockird31/OnlineStore/gen/auth"
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
