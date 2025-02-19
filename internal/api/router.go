package api

import (
	"net/http"

	"github.com/Suhaan-Bhandary/go-api-template/internal/app"
	"github.com/Suhaan-Bhandary/go-api-template/internal/pkg/middleware"
	"github.com/go-chi/chi/v5"
)

func NewRouter(services app.Dependencies) *chi.Mux {
	apiRouter := chi.NewRouter()

	// Middleware
	apiRouter.Use(middleware.RequestId)
	apiRouter.Use(middleware.RequestLogger)
	apiRouter.Use(middleware.Recoverer)

	// Routes
	apiRouter.Mount("/api/v1/users", userRouter(services))
	apiRouter.NotFound(NotFoundHandler)

	return apiRouter
}

// Router for user
func userRouter(services app.Dependencies) http.Handler {
	r := chi.NewRouter()

	r.Get("/", ListUsersPaginatedHandler(services.UserService))
	r.Post("/", CreateUserHandler(services.UserService))
	r.Post("/{userId}/job", CreateUserJobHandler())
	r.Get("/panic", UserPanicHandler())

	return r
}
