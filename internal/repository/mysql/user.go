package repository

import (
	"context"
	"fmt"

	"github.com/Suhaan-Bhandary/go-api-template/internal/pkg/constants"
	customerrors "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/customErrors"
	"github.com/Suhaan-Bhandary/go-api-template/internal/repository"
	"github.com/jmoiron/sqlx"

	"github.com/go-sql-driver/mysql"
)

type userStore struct {
	BaseRepository
}

func NewUserRepo(db *sqlx.DB) repository.UserStorer {
	return &userStore{
		BaseRepository: BaseRepository{DB: db},
	}
}

func (userStore *userStore) CreateUser(ctx context.Context, tx repository.Transaction, user repository.User) error {
	queryExecutor := userStore.initiateQueryExecutor(tx)

	query := `
		INSERT INTO users (id, email, password, created_at, updated_at) 
		VALUES (:id, :email, :password, :created_at, :updated_at)
	`

	query, values, err := queryExecutor.BindNamed(query, user)
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}

	_, err = queryExecutor.Exec(query, values...)
	if err != nil {
		sqlErr, ok := err.(*mysql.MySQLError)
		if ok && sqlErr.Number == constants.MYSQL_KEY_EXITS {
			return customerrors.DuplicateKeyError{Message: "email already exits"}
		}

		return fmt.Errorf("failed to insert user: %w", err)
	}

	return nil
}
