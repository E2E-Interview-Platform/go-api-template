package api

import (
	"net/http"

	customerrors "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/customErrors"
	"github.com/Suhaan-Bhandary/go-api-template/internal/pkg/middleware"
)

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	middleware.ErrorResponse(w, http.StatusNotFound, customerrors.NotFoundError{Message: "API not Found"})
	return
}
