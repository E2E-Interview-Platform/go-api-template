package customerrors

import (
	"net/http"
)

func MapError(err error) (int, error) {
	switch typedError := err.(type) {
	case BadRequestError:
		return http.StatusBadRequest, err
	case DuplicateKeyError:
		return http.StatusBadRequest, err
	case InvalidCredentialError:
		return http.StatusUnauthorized, err
	case NotFoundError:
		return http.StatusNotFound, err
	case CustomError:
		return typedError.Code, err
	default:
		return http.StatusInternalServerError, err
	}
}
