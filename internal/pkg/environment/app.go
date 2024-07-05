package environment

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	customerrors "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/customErrors"
)

var (
	// required: false value: production
	// Denotes the current Environment of the application
	ENVIRONMENT string

	// required: true value: port number
	// Represents the port on which the application is served
	PORT int

	// required: true value: string
	// Used as a secret key for JWT token
	JWT_SECRET_KEY string
)

func SetupAppEnvironment() error {
	var err error

	// ENVIRONMENT
	ENVIRONMENT = os.Getenv("ENVIRONMENT")

	// PORT
	PORT, err = getPORT()
	if err != nil {
		return err
	}

	// JWT_SECRET_KEY
	JWT_SECRET_KEY = os.Getenv("JWT_SECRET_KEY")
	if JWT_SECRET_KEY == "" {
		return customerrors.Error{
			CustomMessage: "Please provide `JWT_SECRET_KEY`",
			InternalError: errors.New("environment variable `JWT_SECRET_KEY` not found"),
		}
	}

	return nil
}

func getPORT() (int, error) {
	strPort := os.Getenv("PORT")
	if strPort == "" {
		return -1, customerrors.Error{
			CustomMessage: "Please provide `PORT`",
			InternalError: errors.New("environment variable `PORT` not found"),
		}
	}

	port, err := strconv.Atoi(strPort)
	if err != nil {
		err = fmt.Errorf("error %w when parsing Environment variable `PORT`", err)
		return -1, customerrors.Error{
			CustomMessage: "Please provide valid value for `PORT`",
			InternalError: err,
		}
	}

	return port, nil
}
