package storage

import (
	"database/sql"
	service "debtsapp/internal/service/model"
	"debtsapp/internal/storage/model"
	log "github.com/sirupsen/logrus"
	"time"
)

type UserStore struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{db: db}
}

func (u *UserStore) Create(user *service.UserRequest) error {

	userEntity := model.UserEntity{
		ID:        3,
		FirstName: user.Name,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: "",
		UpdatedAt: "",
	}

	query := "insert into users (first_name, last_name, username, email, password, created_at, updated_at) values ( $1, $2, $3, $4, $5, $6, $7)"

	_, _ = u.db.Exec(query,
		userEntity.FirstName,
		userEntity.LastName,
		userEntity.Username,
		userEntity.Email,
		userEntity.Password,
		time.Now(),
		time.Now())

	log.Info("User inserted")

	return nil

}
