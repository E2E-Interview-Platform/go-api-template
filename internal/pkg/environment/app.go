package environment

import (
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

	return nil
}

func getPORT() (int, error) {
	strPort := os.Getenv("PORT")
	if strPort == "" {
		return -1, customerrors.CustomError{Message: "environment variable `PORT` not found"}
	}

	port, err := strconv.Atoi(strPort)
	if err != nil {
		return -1, fmt.Errorf("error %w when parsing Environment variable `PORT`", err)
	}

	return port, nil
}
