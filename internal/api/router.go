package api

import (
	"net/http"

	"github.com/Suhaan-Bhandary/go-api-template/internal/pkg/middleware"
	"github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux {
	apiRouter := chi.NewRouter()

	// Middleware
	apiRouter.Use(middleware.Recoverer)

	// Routes
	apiRouter.Mount("/api/v1/users", userRouter())
	apiRouter.NotFound(notFoundHandler)

	return apiRouter
}

// Router for user
func userRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", userDetails())
	r.Get("/panic", userPanic())

	return r
}
