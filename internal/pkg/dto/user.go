package dto

import (
	"context"
	"regexp"

	"github.com/Suhaan-Bhandary/go-api-template/internal/pkg/constants"
	ctxlogger "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/ctxLogger"
	customerrors "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/customErrors"
)

type User struct {
	ID        string `json:"id" db:"id"`
	Email     string `json:"email" db:"email"`
	CreatedAt int64  `json:"created_at" db:"created_at"`
	UpdatedAt int64  `json:"updated_at" db:"updated_at"`
}

// List User
type ListUsersRequest struct {
	Page        int    `json:"page"`
	Limit       int    `json:"limit"`
	SearchValue string `json:"search"`
}

func (req *ListUsersRequest) Validate(ctx context.Context) error {
	if req.Page <= 0 {
		err := customerrors.BadRequestError{Message: "page value should be greater than 0"}
		ctxlogger.Error(ctx, "list user validate err: %s", err.Error())
		return err
	}

	if req.Limit <= 0 || req.Limit > 1000 {
		err := customerrors.BadRequestError{Message: "limit should be between 1 and 1000"}
		ctxlogger.Error(ctx, "list user validate err: %s", err.Error())
		return err
	}

	return nil
}

type ListUsersResponse struct {
	Users []User `json:"users"`
}

// Create User
type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req *CreateUserRequest) Validate(ctx context.Context) error {
	if req.Email == "" {
		err := customerrors.BadRequestError{Message: "email is required"}
		ctxlogger.Error(ctx, "create user validate err: %s", err.Error())
		return err
	}

	// check if email is in correct format
	match, err := regexp.MatchString(constants.EMAIL_REGEX, req.Email)
	if err != nil || !match {
		err := customerrors.BadRequestError{Message: "invalid email"}
		ctxlogger.Error(ctx, "create user validate err: %s", err.Error())
		return err
	}

	if req.Password == "" {
		err := customerrors.BadRequestError{Message: "password required"}
		ctxlogger.Error(ctx, "create user validate err: %s", err.Error())
		return err
	}

	return nil
}

type CreateUserResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}
