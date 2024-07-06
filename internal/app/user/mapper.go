package user

import (
	"context"
	"net/http"
	"time"

	ctxlogger "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/ctxLogger"
	customerrors "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/customErrors"
	"github.com/Suhaan-Bhandary/go-api-template/internal/pkg/dto"
	"github.com/Suhaan-Bhandary/go-api-template/internal/pkg/helpers"
	"github.com/Suhaan-Bhandary/go-api-template/internal/repository"
)

func (userSvc *service) CreateUserRequestToUserMapper(ctx context.Context, userDetail dto.CreateUserRequest) (repository.User, error) {
	ctxlogger.Info(ctx, "Starting #CreateUserRequestToUserMapper")
	defer ctxlogger.Info(ctx, "Ending #CreateUserRequestToUserMapper")

	hashedPassword, err := helpers.Hash(userDetail.Password)
	if err != nil {
		ctxlogger.Error(ctx, "error hashing password in #CreateUserRequestToUserMapper")
		return repository.User{}, customerrors.Error{
			Code:          http.StatusInternalServerError,
			CustomMessage: "Something went wrong",
			InternalError: err,
		}
	}

	currentUnixMilliSeconds := time.Now().UnixMilli()
	return repository.User{
		ID:        userSvc.idGenerator.GenerateId(),
		Email:     userDetail.Email,
		Password:  hashedPassword,
		CreatedAt: currentUnixMilliSeconds,
		UpdatedAt: currentUnixMilliSeconds,
	}, nil
}
