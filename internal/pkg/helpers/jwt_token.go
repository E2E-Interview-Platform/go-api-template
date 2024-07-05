package helpers

import (
	"fmt"
	"net/http"
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
		return nil, customerrors.Error{
			Code:          http.StatusUnauthorized,
			CustomMessage: "JWT token is invalid",
			InternalError: fmt.Errorf("JWT verification error, err: %s", err),
		}
	}

	return token, nil
}
