package helpers

import (
	"fmt"
	"time"

	"github.com/Suhaan-Bhandary/go-api-template/internal/pkg/constants"
	customerrors "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/customErrors"
	"github.com/Suhaan-Bhandary/go-api-template/internal/pkg/environment"
	"github.com/golang-jwt/jwt/v5"
)

type TokenDetails struct {
	ID string `json:"id"`
}

func GenerateToken(tokenDetails TokenDetails) (string, error) {
	expirationTime := time.Now().Add(time.Hour * constants.JWT_EXPIRATION_TIME_HOURS)

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":  tokenDetails.ID,
			"exp": expirationTime.Unix(),
		},
	)

	tokenString, err := token.SignedString([]byte(environment.JWT_SECRET_KEY))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(environment.JWT_SECRET_KEY), nil
	})

	if err != nil || !token.Valid {
		fmt.Println(err, token.Valid)
		return nil, customerrors.InvalidCredentialError{Message: "invalid JWT token"}
	}

	return token, nil
}
