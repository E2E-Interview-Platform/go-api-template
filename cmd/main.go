package main

import (
	"net/http"

	customerrors "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/customErrors"
	customMiddleware "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	apiRouter := chi.NewRouter()

	// Application Middleware
	apiRouter.Use(middleware.RequestID)

	// Creating Sub routers
	apiRouter.Mount("/api/v1/users", userRouter())

	// Not Found Route
	apiRouter.NotFound(func(w http.ResponseWriter, r *http.Request) {
		customMiddleware.ErrorResponse(w, http.StatusNotFound, customerrors.NotFoundError{Message: "API not Found"})
		return
	})

	http.ListenAndServe(":8080", apiRouter)
}

// Router for user
func userRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		customMiddleware.SuccessResponse(w, http.StatusOK, struct{ Name string }{Name: "suhaan"})
		return
	})

	return r
}
