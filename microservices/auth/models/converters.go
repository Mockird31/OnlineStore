package model

import (
	authProto "github.com/Mockird31/OnlineStore/gen/auth"
)

func UserIDToInt(userID *authProto.UserID) int64 {
	return userID.Id
}

func IntToUserID(userID int64) *authProto.UserID {
	return &authProto.UserID{Id: userID}
}

func StringToSessionID(sessionID string) *authProto.SessionID {
	return &authProto.SessionID{Id: sessionID}
}

func SessionIDToString(sessionID *authProto.SessionID) string {
	return sessionID.Id
}