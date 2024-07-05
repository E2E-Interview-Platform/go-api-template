package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/Suhaan-Bhandary/go-api-template/internal/pkg/context"
	ctxlogger "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/ctxLogger"
	customerrors "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/customErrors"
	"github.com/google/uuid"
)

func RequestId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		requestID := r.Header.Get("X-Request-ID")
		if requestID == "" {
			requestID = uuid.NewString()
		}

		ctx = context.SetRequestID(ctx, requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func Recoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		defer func() {
			if rvr := recover(); rvr != nil {
				// Ignoring ErrAbortHandler
				if rvr == http.ErrAbortHandler {
					panic(rvr)
				}

				ctxlogger.Info(ctx, "panic: %v, stack: %s", rvr, string(debug.Stack()))

				ErrorResponse(ctx, w, ErrorResponseOptions{
					Error: customerrors.Error{
						Code:          http.StatusInternalServerError,
						CustomMessage: "Something went wrong",
						InternalError: fmt.Errorf("recovering panic: %s", rvr),
					},
				})
			}
		}()

		next.ServeHTTP(w, r)
	})
}
