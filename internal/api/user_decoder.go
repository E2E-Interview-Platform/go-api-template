package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	ctxlogger "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/ctxLogger"
	customerrors "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/customErrors"
	"github.com/Suhaan-Bhandary/go-api-template/internal/pkg/dto"
)

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
	req.Password = strings.TrimSpace(req.Password)

	return req, nil
}
