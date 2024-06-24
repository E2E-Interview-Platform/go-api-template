package middleware

import (
	"fmt"
	"net/http"

	"github.com/Suhaan-Bhandary/go-api-template/internal/pkg/context"
)

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rid := context.GetRequestID(r.Context())

		fmt.Printf("[rid=%s] %s %s Request Started\n", rid, r.Method, r.URL.Path)
		defer fmt.Printf("[rid=%s] %s %s Request Ended\n", rid, r.Method, r.URL.Path)

		next.ServeHTTP(w, r)
	})
}
