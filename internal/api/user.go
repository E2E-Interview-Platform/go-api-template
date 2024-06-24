package api

import (
	"fmt"
	"net/http"

	"github.com/Suhaan-Bhandary/go-api-template/internal/pkg/middleware"
)

func userDetails() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("User Details Function")
		middleware.SuccessResponse(w, http.StatusOK, struct{ Name string }{Name: "suhaan"})
		return
	}
}

func userPanic() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		panic("Trying Panic")
	}
}
