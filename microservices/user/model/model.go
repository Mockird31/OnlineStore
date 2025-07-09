package model

type RegisterData struct {
	Username        string
	Email           string
	Password        string
	ConfirmPassword string
}

type User struct {
	Id       int64
	Username string
	Email    string
}
