package environment

import (
	"fmt"
	"os"

	customerrors "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/customErrors"
)

var (
	// required: true value: "root"
	// Username of the db
	DB_USER string

	// required: true value: "password"
	// Password of the db
	DB_PASSWORD string

	// required: true value: "domain:port"
	// URL of database
	DB_URL string

	// required: true value: "database_name"
	// name of database
	DB_NAME string
)

func SetupDatabaseEnvironment() error {
	DB_USER = os.Getenv("DB_USER")
	if DB_USER == "" {
		return customerrors.Error{
			CustomMessage: "Please provide `DB_USER`",
			InternalError: fmt.Errorf("environment variable `DB_USER` not found"),
		}
	}

	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	if DB_PASSWORD == "" {
		return customerrors.Error{
			CustomMessage: "Please provide `DB_PASSWORD`",
			InternalError: fmt.Errorf("environment variable `DB_PASSWORD` not found"),
		}
	}

	DB_URL = os.Getenv("DB_URL")
	if DB_URL == "" {
		return customerrors.Error{
			CustomMessage: "Please provide `DB_URL`",
			InternalError: fmt.Errorf("environment variable `DB_URL` not found"),
		}
	}

	DB_NAME = os.Getenv("DB_NAME")
	if DB_NAME == "" {
		return customerrors.Error{
			CustomMessage: "Please provide `DB_NAME`",
			InternalError: fmt.Errorf("environment variable `DB_NAME` not found"),
		}
	}

	return nil
}
