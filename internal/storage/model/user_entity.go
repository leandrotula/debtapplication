package model

type UserEntity struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"_"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
