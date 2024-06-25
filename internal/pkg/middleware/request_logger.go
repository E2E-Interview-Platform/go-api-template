package middleware

import (
	"net/http"

	ctxlogger "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/ctxLogger"
)

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ctxlogger.Info(ctx, "%s %s Request Started\n", r.Method, r.URL.Path)
		defer ctxlogger.Info(ctx, "%s %s Request Ended\n", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)
	})
}
