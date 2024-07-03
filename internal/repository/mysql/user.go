package repository

import (
	"context"
	"fmt"

	"github.com/Suhaan-Bhandary/go-api-template/internal/pkg/constant"
	customerrors "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/customErrors"
	"github.com/Suhaan-Bhandary/go-api-template/internal/repository"
	"github.com/jmoiron/sqlx"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
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
		if err.(*mysql.MySQLError).Number == constant.MYSQL_KEY_EXITS {
			return customerrors.DuplicateKeyError{Message: "email already exits"}
		}

		return fmt.Errorf("failed to insert user: %w", err)
	}

	return nil
}
