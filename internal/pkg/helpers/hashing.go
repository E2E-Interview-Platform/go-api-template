package helpers

import (
	customerrors "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/customErrors"
	"golang.org/x/crypto/bcrypt"
)

func Hash(data string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data), 10)
	if err != nil {
		return "", customerrors.InternalServerError{
			Message: "internal server error",
		}
	}

	return string(hashedPassword), nil
}
