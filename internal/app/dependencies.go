package app

import (
	"github.com/Suhaan-Bhandary/go-api-template/internal/app/user"
	repository "github.com/Suhaan-Bhandary/go-api-template/internal/repository/mysql"
	"github.com/jmoiron/sqlx"
)

type Dependencies struct {
	UserService user.Service
}

func NewServices(db *sqlx.DB) Dependencies {
	// Repository
	userRepo := repository.NewUserRepo(db)

	// Services
	userService := user.NewService(userRepo)

	return Dependencies{
		UserService: userService,
	}
}
