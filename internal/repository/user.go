package repository

import (
	"context"

	"github.com/Suhaan-Bhandary/go-api-template/internal/pkg/dto"
)

type UserStorer interface {
	RepositoryTransaction

	ListUsersPaginated(ctx context.Context, tx Transaction, filters dto.ListUsersPaginatedRequest) (dto.PaginatedUsers, error)
	CreateUser(ctx context.Context, tx Transaction, user User) error
}

type User struct {
	ID        string `json:"id" db:"id"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"password"`
	CreatedAt int64  `json:"created_at" db:"created_at"`
	UpdatedAt int64  `json:"updated_at" db:"updated_at"`
}
