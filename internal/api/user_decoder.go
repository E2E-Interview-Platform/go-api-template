package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	ctxlogger "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/ctxLogger"
	customerrors "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/customErrors"
	"github.com/Suhaan-Bhandary/go-api-template/internal/pkg/dto"
)

func decodeListUsersRequest(ctx context.Context, r *http.Request) (dto.ListUsersPaginatedRequest, error) {
	// Getting data from search query
	query := r.URL.Query()

	search := query.Get("search")

	page, err := strconv.Atoi(query.Get("page"))
	if err != nil {
		err = fmt.Errorf("error getting page from URL, err: %s", err)
		ctxlogger.Error(ctx, err.Error())
		return dto.ListUsersPaginatedRequest{}, customerrors.Error{
			Code:          http.StatusBadRequest,
			CustomMessage: "Please provide page value",
			InternalError: err,
		}
	}

	limit, err := strconv.Atoi(query.Get("limit"))
	if err != nil {
		ctxlogger.Error(ctx, "error getting limit from URL, err: %s", err)
		return dto.ListUsersPaginatedRequest{}, customerrors.Error{
			Code:          http.StatusBadRequest,
			CustomMessage: "Please provide limit value",
			InternalError: err,
		}
	}

	req := dto.ListUsersPaginatedRequest{
		SearchValue: strings.TrimSpace(search),
		Page:        page,
		Limit:       limit,
	}

	return req, nil
}

func decodeCreateUserRequest(ctx context.Context, r *http.Request) (dto.CreateUserRequest, error) {
	var req dto.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		err = fmt.Errorf("error decoding create user request, err: %s", err)
		ctxlogger.Error(ctx, err.Error())
		return dto.CreateUserRequest{}, customerrors.Error{
			Code:          http.StatusBadRequest,
			CustomMessage: "Please provide a valid request body",
			InternalError: err,
		}
	}

	// Removing white spaces
	req.Email = strings.TrimSpace(req.Email)

	return req, nil
}
