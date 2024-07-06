package api

import (
	"net/http"

	"github.com/Suhaan-Bhandary/go-api-template/internal/app/user"
	ctxlogger "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/ctxLogger"
	"github.com/Suhaan-Bhandary/go-api-template/internal/pkg/dto"
	"github.com/Suhaan-Bhandary/go-api-template/internal/pkg/middleware"
)

func ListUsersPaginatedHandler(userSvc user.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctxlogger.Info(ctx, "Starting List Users Handler")
		defer ctxlogger.Info(ctx, "Ending List Users Handler")

		req, err := decodeListUsersRequest(ctx, r)
		if err != nil {
			middleware.ErrorResponse(ctx, w, middleware.ErrorResponseOptions{Error: err})
			return
		}

		err = req.Validate(ctx)
		if err != nil {
			middleware.ErrorResponse(ctx, w, middleware.ErrorResponseOptions{Error: err})
			return
		}

		paginatedUsers, err := userSvc.ListUsersPaginated(ctx, req)
		if err != nil {
			middleware.ErrorResponse(ctx, w, middleware.ErrorResponseOptions{Error: err})
			return
		}

		// User and pagination data
		users := paginatedUsers.Users
		pagination := paginatedUsers.Pagination

		middleware.SuccessResponse(ctx, w, http.StatusOK, dto.ListUsersPaginatedResponse{
			Users:      users,
			Pagination: pagination,
		})
	}
}

func CreateUserHandler(userSvc user.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctxlogger.Info(ctx, "Starting Create User Handler")
		defer ctxlogger.Info(ctx, "Ending Create User Handler")

		req, err := decodeCreateUserRequest(ctx, r)
		if err != nil {
			middleware.ErrorResponse(ctx, w, middleware.ErrorResponseOptions{Error: err})
			return
		}

		err = req.Validate(ctx)
		if err != nil {
			middleware.ErrorResponse(ctx, w, middleware.ErrorResponseOptions{Error: err})
			return
		}

		token, err := userSvc.CreateUser(ctx, req)
		if err != nil {
			middleware.ErrorResponse(ctx, w, middleware.ErrorResponseOptions{Error: err})
			return
		}

		middleware.SuccessResponse(ctx, w, http.StatusOK, dto.CreateUserResponse{
			Message: "User created successfully",
			Token:   token,
		})
	}
}

func UserPanicHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctxlogger.Info(ctx, "Starting User Panic")
		defer ctxlogger.Info(ctx, "Ending User Panic")

		panic("Trying Panic")
	}
}
