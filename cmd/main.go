package main

import (
	"context"
	"net/http"

	"github.com/Suhaan-Bhandary/go-api-template/internal/api"
	ctxlogger "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/ctxLogger"
	"github.com/Suhaan-Bhandary/go-api-template/internal/pkg/environment"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "rid", "main-function")

	err := environment.LoadEnvironment()
	if err != nil {
		ctxlogger.Error(ctx, err.Error())
		return
	}

	apiRouter := api.NewRouter()

	err = http.ListenAndServe(":8080", apiRouter)
	if err != nil {
		ctxlogger.Error(ctx, err.Error())
	}
}
