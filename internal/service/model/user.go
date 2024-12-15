package model

type UserRequest struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Username string `json:"username"`
	Password string `json:"-"`
	Email    string `json:"email"`
}
