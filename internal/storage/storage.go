package storage

import (
	"context"
	"database/sql"
	"debtsapp/internal/env"
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

type ErrorDB struct {
	Message string
}

type Storage struct {
	Users interface {
		Create(user *User) error
	}
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		Users: NewUserStore(db),
	}
}

func NewErrorDB(message string) *ErrorDB {
	return &ErrorDB{
		Message: message,
	}
}

func New() (*sql.DB, error) {
	user := env.GetString("POSTGRES_USER", "postgres")
	dbname := env.GetString("POSTGRES_DB", "debts")
	sslMode := env.GetString("POSTGRES_SSLMODE", "disable")
	password := env.GetString("POSTGRES_PASSWORD", "debtspassword")
	host := env.GetString("POSTGRES_HOST", "localhost")
	maxOpenConnections := env.GetInt("POSTGRES_MAX_OPEN_CONNECTIONS", 10)
	maxIdleConnections := env.GetInt("POSTGRES_MAX_IDLE_CONNECTIONS", 5)
	maxIdleTimeConnection := env.GetString("POSTGRES_MAX_IDLE_TIME_CONNECTION", "10m")

	dsn := fmt.Sprintf("user=%s dbname=%s sslMode=%s password=%s host=%s", user, dbname, sslMode, password, host)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxIdleConns(maxIdleConnections)
	db.SetMaxOpenConns(maxOpenConnections)

	duration, err := time.ParseDuration(maxIdleTimeConnection)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(duration)

	// handle time out against db
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}