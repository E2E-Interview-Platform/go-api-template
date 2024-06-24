package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/Suhaan-Bhandary/go-api-template/internal/pkg/context"
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
		defer func() {
			if rvr := recover(); rvr != nil {
				// Ignoring ErrAbortHandler
				if rvr == http.ErrAbortHandler {
					panic(rvr)
				}

				rid := context.GetRequestID(r.Context())
				fmt.Printf("[rid=%s] panic: %v, stack: %s", rid, rvr, string(debug.Stack()))

				ErrorResponse(w, http.StatusInternalServerError, errors.New("Internal Server Error"))
			}
		}()

		next.ServeHTTP(w, r)
	})
}
