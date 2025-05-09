package model

type UserRequest struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserResponse struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type RegisterUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func NewUserResponse(name, lastName, username, email string) *UserResponse {
	return &UserResponse{
		Name:     name,
		LastName: lastName,
		Username: username,
		Email:    email,
	}
}
