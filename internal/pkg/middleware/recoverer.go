package middleware

import (
	"errors"
	"net/http"
	"runtime/debug"

	"github.com/Suhaan-Bhandary/go-api-template/internal/pkg/context"
	ctxlogger "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/ctxLogger"
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

				ErrorResponse(ctx, w, http.StatusInternalServerError, errors.New("Internal Server Error"))
			}
		}()

		next.ServeHTTP(w, r)
	})
}
