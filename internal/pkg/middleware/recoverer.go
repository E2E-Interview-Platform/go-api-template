package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"
)

func Recoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil {
				// Ignoring ErrAbortHandler
				if rvr == http.ErrAbortHandler {
					panic(rvr)
				}

				fmt.Printf("panic: %v, stack: %s", r, string(debug.Stack()))
				ErrorResponse(w, http.StatusInternalServerError, errors.New("Internal Server Error"))
			}
		}()

		next.ServeHTTP(w, r)
	})
}
