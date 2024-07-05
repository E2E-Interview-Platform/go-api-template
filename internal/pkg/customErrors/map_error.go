package customerrors

import "net/http"

func MapError(err error) (int, string, string) {
	switch val := err.(type) {
	case Error:
		return val.Code, val.CustomMessage, val.InternalError.Error()
	default:
		return http.StatusInternalServerError, "Something went wrong", val.Error()
	}
}
