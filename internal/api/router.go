package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux {
	apiRouter := chi.NewRouter()

	// Routes
	apiRouter.Mount("/api/v1/users", userRouter())
	apiRouter.NotFound(notFoundHandler)

	return apiRouter
}

// Router for user
func userRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", userDetails())

	return r
}
