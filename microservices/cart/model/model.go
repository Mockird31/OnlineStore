package model

type CartIdentifier interface {
	IsCartIdentifier() int64
}

type UserID struct {
	UserID int64
}

func (ui *UserID) IsCartIdentifier() int64 {
	return 1
}

type SessionID struct {
	SessionID string
}

func (si *SessionID) IsCartIdentifier() int64 {
	return -1
}

type AddRequest struct {
	CartID CartIdentifier
	ItemID int64
}

type DeleteRequest struct {
	CartID CartIdentifier
	ItemID int64
}

type Cart struct {
	CartID   CartIdentifier
	ItemsIDs []int64
}
