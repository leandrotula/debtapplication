package storage

import (
	"context"
	"database/sql"
	"fmt"
)

type PostStore struct {
	db *sql.DB
}

func NewPostStore(db *sql.DB) *PostStore {
	return &PostStore{db: db}
}

func (s *PostStore) Create(ctx context.Context) error {
	fmt.Print("called here")
	return nil
}
