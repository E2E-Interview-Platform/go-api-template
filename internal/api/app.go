package api

import (
	"net/http"

	customerrors "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/customErrors"
	"github.com/Suhaan-Bhandary/go-api-template/internal/pkg/middleware"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	middleware.ErrorResponse(ctx, w, middleware.ErrorResponseOptions{Error: customerrors.NotFoundError{Message: "API not Found"}})
	return
}
