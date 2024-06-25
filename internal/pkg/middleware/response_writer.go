package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	ctxlogger "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/ctxLogger"
)

type response struct {
	ErrorCode    int         `json:"error_code"`
	ErrorMessage string      `json:"error_message"`
	Data         interface{} `json:"data"`
}

func SuccessResponse(ctx context.Context, w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	payload := response{
		Data: data,
	}

	out, err := json.Marshal(payload)
	if err != nil {
		ctxlogger.Error(ctx, "cannot marshal success response payload")
		writeServerErrorResponse(ctx, w)
		return
	}

	_, err = w.Write(out)
	if err != nil {
		ctxlogger.Error(ctx, "cannot write json success response")
		writeServerErrorResponse(ctx, w)
		return
	}
}

func ErrorResponse(ctx context.Context, w http.ResponseWriter, httpStatus int, err error) {
	// Printing the error
	ctxlogger.Error(ctx, err.Error())

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)

	payload := response{
		ErrorCode:    httpStatus,
		ErrorMessage: err.Error(),
	}

	out, err := json.Marshal(payload)
	if err != nil {
		ctxlogger.Error(ctx, "error occurred while marshaling response payload")
		writeServerErrorResponse(ctx, w)
		return
	}

	_, err = w.Write(out)
	if err != nil {
		ctxlogger.Error(ctx, "error occurred while writing response")
		writeServerErrorResponse(ctx, w)
		return
	}
}

func writeServerErrorResponse(ctx context.Context, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	_, err := w.Write([]byte(fmt.Sprintf("{\"message\":%s}", "internal server error")))
	if err != nil {
		ctxlogger.Error(ctx, "error occurred while writing response")
	}
}
