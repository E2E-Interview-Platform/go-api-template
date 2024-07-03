package repository

import (
	"context"
	"fmt"

	"github.com/Suhaan-Bhandary/go-api-template/internal/repository"
	"github.com/jmoiron/sqlx"
)

type userStore struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) repository.UserStorer {
	return &userStore{
		db: db,
	}
}

func (userStore *userStore) CreateUser(ctx context.Context, user repository.User) error {
	query := `
		INSERT INTO users (id, email, password, created_at, updated_at) 
		VALUES (:id, :email, :password, :created_at, :updated_at)
	`

	_, err := userStore.db.NamedExec(query, user)
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}

	return nil
}
