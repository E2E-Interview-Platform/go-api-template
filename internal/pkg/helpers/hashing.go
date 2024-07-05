package helpers

import (
	"fmt"
	"net/http"

	customerrors "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/customErrors"
	"golang.org/x/crypto/bcrypt"
)

func Hash(data string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data), 10)
	if err != nil {
		return "", customerrors.Error{
			Code:          http.StatusInternalServerError,
			CustomMessage: "Something went wrong",
			InternalError: fmt.Errorf("error hashing password, err: %s", err),
		}
	}

	return string(hashedPassword), nil
}
