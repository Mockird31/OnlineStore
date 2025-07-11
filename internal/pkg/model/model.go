package model

type RegisterData struct {
	Username        string `validate:"min=5,max=30"`
	Email           string
	Password        string
	ConfirmPassword string
}

type User struct {
	Id       int64
	Username string
	Email    string
}

type APIResponse struct {
	Status int
	Body   interface{}
}
