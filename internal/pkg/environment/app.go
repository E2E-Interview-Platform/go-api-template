package environment

import "os"

var (
	// required: false value: production
	// Denotes the current Environment of the application
	ENVIRONMENT string
)

func SetupAppEnvironment() error {
	ENVIRONMENT = os.Getenv("ENVIRONMENT")

	return nil
}
