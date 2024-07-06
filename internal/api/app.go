package api

import (
	"fmt"
	"net/http"

	"github.com/Suhaan-Bhandary/go-api-template/internal/pkg/context"
	ctxlogger "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/ctxLogger"
	customerrors "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/customErrors"
	"github.com/Suhaan-Bhandary/go-api-template/internal/pkg/middleware"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	rid := context.GetRequestID(ctx)

	ctxlogger.Info(ctx, "Starting not found handler")
	defer ctxlogger.Info(ctx, "Ending not found handler")

	middleware.ErrorResponse(ctx, w,
		middleware.ErrorResponseOptions{
			Error: customerrors.Error{
				Code:          http.StatusNotFound,
				CustomMessage: "API not found",
				InternalError: fmt.Errorf("api not found, rid=%s", rid),
			},
		},
	)
}
