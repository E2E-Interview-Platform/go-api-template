package dto

import (
	"context"
	"regexp"

	"github.com/Suhaan-Bhandary/go-api-template/internal/pkg/constant"
	ctxlogger "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/ctxLogger"
	customerrors "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/customErrors"
)

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
	match, err := regexp.MatchString(constant.EMAIL_REGEX, req.Email)
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
