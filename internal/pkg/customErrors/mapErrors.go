package customerrors

import "net/http"

func MapError(err error) (int, error) {
	switch typedError := err.(type) {
	case NotFoundError:
		return http.StatusNotFound, err
	case BadRequestError:
		return http.StatusBadRequest, err
	case CustomError:
		return typedError.Code, err
	default:
		return http.StatusInternalServerError, err
	}
}
