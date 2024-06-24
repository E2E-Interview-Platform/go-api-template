package main

import (
	"net/http"

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
		w.Write([]byte("Not Found"))
	})

	http.ListenAndServe(":8080", apiRouter)
}

// Router for user
func userRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("/api/v1/user router"))
	})

	return r
}
