package dto

import (
	"context"
	"errors"
	"net/http"
	"regexp"

	"github.com/Suhaan-Bhandary/go-api-template/internal/pkg/constants"
	customerrors "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/customErrors"
)

type User struct {
	ID        string `json:"id" db:"id"`
	Email     string `json:"email" db:"email"`
	CreatedAt int64  `json:"created_at" db:"created_at"`
	UpdatedAt int64  `json:"updated_at" db:"updated_at"`
}

type PaginatedUsers struct {
	Users      []User     `json:"users"`
	Pagination Pagination `json:"pagination"`
}

// List User
type ListUsersRequest struct {
	Page        int    `json:"page"`
	Limit       int    `json:"limit"`
	SearchValue string `json:"search"`
}

func (req *ListUsersRequest) Validate(ctx context.Context) error {
	if req.Page <= 0 {
		return customerrors.Error{
			Code:          http.StatusBadRequest,
			CustomMessage: "Please provide value greater than 0",
			InternalError: errors.New("list user validate err: page value should be greater than 0"),
		}
	}

	if req.Limit <= 0 || req.Limit > 1000 {
		return customerrors.Error{
			Code:          http.StatusBadRequest,
			CustomMessage: "Limit should be between 1 and 1000",
			InternalError: errors.New("list user validate err: limit should be between 1 and 1000"),
		}
	}

	return nil
}

type ListUsersResponse PaginatedUsers

// Create User
type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req *CreateUserRequest) Validate(ctx context.Context) error {
	if req.Email == "" {
		return customerrors.Error{
			Code:          http.StatusBadRequest,
			CustomMessage: "Please provide an email",
			InternalError: errors.New("create user validate err: email is required"),
		}
	}

	// check if email is in correct format
	match, err := regexp.MatchString(constants.EMAIL_REGEX, req.Email)
	if err != nil || !match {
		return customerrors.Error{
			Code:          http.StatusBadRequest,
			CustomMessage: "Please provide a valid email",
			InternalError: errors.New("create user validate err: email is invalid"),
		}
	}

	if req.Password == "" {
		return customerrors.Error{
			Code:          http.StatusBadRequest,
			CustomMessage: "Please provide a password",
			InternalError: errors.New("create user validate err: password is required"),
		}
	}

	return nil
}

type CreateUserResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}
