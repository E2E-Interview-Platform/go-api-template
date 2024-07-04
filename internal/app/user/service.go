package user

import (
	"context"

	ctxlogger "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/ctxLogger"
	"github.com/Suhaan-Bhandary/go-api-template/internal/pkg/dto"
	"github.com/Suhaan-Bhandary/go-api-template/internal/pkg/helpers"
	"github.com/Suhaan-Bhandary/go-api-template/internal/repository"
)

type service struct {
	userRepo    repository.UserStorer
	idGenerator helpers.IDGenerator
}

type Service interface {
	CreateUser(ctx context.Context, userDetail dto.CreateUserRequest) (string, error)
}

func NewService(userRepo repository.UserStorer) Service {
	return &service{
		userRepo:    userRepo,
		idGenerator: helpers.UUIDGenerator{},
	}
}

func (userSvc *service) CreateUser(ctx context.Context, userDetail dto.CreateUserRequest) (string, error) {
	ctxlogger.Info(ctx, "Create User Service")

	tx, err := userSvc.userRepo.BeginTx(ctx)
	if err != nil {
		ctxlogger.Info(ctx, "error in creating transaction, err: %s", err.Error())
		return "", err
	}
	defer func() {
		rvr := recover()
		defer func() {
			if rvr != nil {
				ctxlogger.Info(ctx, "Transaction aborted because of panic: %v, Propagating panic further", rvr)
				panic(rvr)
			}
		}()

		txErr := userSvc.userRepo.HandleTransaction(ctx, tx, err == nil && rvr == nil)
		if txErr != nil {
			err = txErr
			ctxlogger.Info(ctx, "error in creating transaction, err: %s", txErr.Error())
			return
		}
	}()

	user, err := userSvc.CreateUserRequestToUserMapper(ctx, userDetail)
	if err != nil {
		ctxlogger.Info(ctx, "error mapping user in create user, err: %s", err.Error())
		return "", err
	}

	err = userSvc.userRepo.CreateUser(ctx, tx, user)
	if err != nil {
		ctxlogger.Info(ctx, "error creating user, err: %s", err.Error())
		return "", err
	}

	token, err := helpers.GenerateToken(helpers.TokenDetails{
		ID: user.ID,
	})
	if err != nil {
		ctxlogger.Info(ctx, "error generating token in create user, err: %s", err.Error())
		return "", err
	}

	return token, nil
}
