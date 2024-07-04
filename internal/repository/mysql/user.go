package repository

import (
	"context"
	"fmt"

	"github.com/Suhaan-Bhandary/go-api-template/internal/pkg/constants"
	customerrors "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/customErrors"
	"github.com/Suhaan-Bhandary/go-api-template/internal/pkg/dto"
	"github.com/Suhaan-Bhandary/go-api-template/internal/repository"
	"github.com/jmoiron/sqlx"

	sq "github.com/Masterminds/squirrel"
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

func (userStore *userStore) ListUsers(ctx context.Context, tx repository.Transaction, filters dto.ListUsersRequest) ([]dto.User, error) {
	queryExecutor := userStore.initiateQueryExecutor(tx)

	queryBuilder := sq.
		Select("id", "email", "created_at", "updated_at").
		From("users")

	// Pagination
	offset := (filters.Page - 1) * filters.Limit
	queryBuilder = queryBuilder.Limit(uint64(filters.Limit)).Offset(uint64(offset))

	// Filtering
	if filters.SearchValue != "" {
		queryBuilder = queryBuilder.Where(sq.Like{"email": "%" + filters.SearchValue + "%"})
	}

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL query: %w", err)
	}

	var users []dto.User
	err = sqlx.Select(queryExecutor, &users, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	return users, nil
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
