package storage

import (
	"context"
	"database/sql"
	service "debtsapp/internal/service/model"
	"debtsapp/internal/storage/model"
	"time"
)

type UserStore struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{db: db}
}

func (u *UserStore) create(ctx context.Context, tx *sql.Tx, user *service.UserRequest) error {

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

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second) // put this as a configuration
	defer cancel()

	query := "insert into users (first_name, last_name, username, email, password, created_at, updated_at) values ( $1, $2, $3, $4, $5, $6, $7)"

	_, _ = tx.Exec(query,
		userEntity.FirstName,
		userEntity.LastName,
		userEntity.Username,
		userEntity.Email,
		userEntity.Password,
		time.Now(),
		time.Now())

	return nil

}

func (u *UserStore) CreateAndInvite(ctx context.Context, user *service.UserRequest, token string) error {
	return withTx(u.db, ctx, func(tx *sql.Tx) error {
		if err := u.create(ctx, tx, user); err != nil {
			return err
		}

		if err := u.createUserInvitation(ctx, tx, user.Email, token); err != nil {
			return err
		}
		return nil
	})
}

func (u *UserStore) createUserInvitation(ctx context.Context, tx *sql.Tx, email string, token string) error {
	query := "insert into user_invitations (token, email) values ($1, $2)"

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second) // put this as a configuration
	defer cancel()

	_, err := tx.ExecContext(ctx, query, token, email)
	if err != nil {
		return err
	}

	return nil

}

func (u *UserStore) Activate(ctx context.Context, token string) error {
	query := "update users set active = true where email = (select email from user_invitations where token = $1)"
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return withTx(u.db, ctx, func(tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, query, token)
		if err != nil {
			return err
		}
		return nil
	})
}
