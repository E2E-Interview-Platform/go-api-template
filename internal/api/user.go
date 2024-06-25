package api

import (
	"net/http"

	ctxlogger "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/ctxLogger"
	"github.com/Suhaan-Bhandary/go-api-template/internal/pkg/middleware"
)

func userDetails() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctxlogger.Info(ctx, "User Details Function")
		middleware.SuccessResponse(ctx, w, http.StatusOK, struct{ Name string }{Name: "suhaan"})
		return
	}
}

func userPanic() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		panic("Trying Panic")
	}
}
