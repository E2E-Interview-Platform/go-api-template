package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	ctxlogger "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/ctxLogger"
	customerrors "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/customErrors"
)

type response struct {
	ErrorCode    int         `json:"error_code"`
	ErrorMessage string      `json:"error_message"`
	Data         interface{} `json:"data"`
}

type ErrorResponseOptions struct {
	HttpStatus   *int
	ErrorMessage *string
	Error        error
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

// Returns error response in json, custom status and message have higher priority than error
func ErrorResponse(ctx context.Context, w http.ResponseWriter, options ErrorResponseOptions) {
	var err error
	var httpStatus int
	var errMessage string

	if options.Error != nil {
		httpStatus, err = customerrors.MapError(options.Error)
		errMessage = err.Error()
	}

	if options.ErrorMessage != nil {
		errMessage = *options.ErrorMessage
	}

	if options.HttpStatus != nil {
		httpStatus = *options.HttpStatus
	}

	// Displaying error
	ctxlogger.Error(ctx, "error response: %s", errMessage)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)

	payload := response{
		ErrorCode:    httpStatus,
		ErrorMessage: errMessage,
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
