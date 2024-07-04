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

func decodeListUsersRequest(ctx context.Context, r *http.Request) (dto.ListUsersRequest, error) {
	// Getting data from search query
	query := r.URL.Query()

	search := query.Get("search")

	page, err := strconv.Atoi(query.Get("page"))
	if err != nil {
		ctxlogger.Error(ctx, "error getting page from URL, err: %s", err)
		return dto.ListUsersRequest{}, customerrors.BadRequestError{Message: "page value not found"}
	}

	limit, err := strconv.Atoi(query.Get("limit"))
	if err != nil {
		ctxlogger.Error(ctx, "error getting limit from URL, err: %s", err)
		return dto.ListUsersRequest{}, customerrors.BadRequestError{Message: "limit value not found"}
	}

	req := dto.ListUsersRequest{
		SearchValue: strings.TrimSpace(search),
		Page:        page,
		Limit:       limit,
	}

	return req, nil
}

func decodeRegisterUserRequest(ctx context.Context, r *http.Request) (dto.CreateUserRequest, error) {
	var req dto.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		err = fmt.Errorf("error decoding register user request, err: %s", err)
		ctxlogger.Error(ctx, err.Error())
		return dto.CreateUserRequest{}, customerrors.BadRequestError{Message: "invalid request body"}
	}

	// Removing white spaces
	req.Email = strings.TrimSpace(req.Email)

	return req, nil
}
