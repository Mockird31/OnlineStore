package model

type RegisterData struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type LoginData struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	Id       int64  `json:"-"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type APIResponse struct {
	Status int         `json:"status" example:"200" description:"HTTP status code"`
	Body   interface{} `json:"body" description:"Response data"`
}
