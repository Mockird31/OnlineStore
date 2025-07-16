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

type Category struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}

type Pagination struct {
	Offset int `sql:"offset"`
	Limit  int `sql:"limit"`
}

type Item struct {
	Id          int64       `json:"id" sql:"id"`
	Title       string      `json:"title" sql:"title"`
	Description string      `json:"description,omitempty" sql:"description"`
	Price       float64     `json:"price" sql:"price"`
	ImageURL    string      `json:"image_url" sql:"image_url"`
	Count       int64       `json:"count" sql:"count"`
	IsActive    bool        `json:"is_active" sql:"is_active"`
	Categories  []*Category `json:"categories" sql:"categories"`
}
