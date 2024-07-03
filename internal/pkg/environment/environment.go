package environment

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvironment() error {
	// Loading .env file if present and ignore if not present
	err := godotenv.Load()
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	err = SetupAppEnvironment()
	if err != nil {
		return err
	}

	err = SetupDatabaseEnvironment()
	if err != nil {
		return err
	}

	return nil
}
