package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Suhaan-Bhandary/go-api-template/internal/api"
	customcontext "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/context"
	ctxlogger "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/ctxLogger"
	"github.com/Suhaan-Bhandary/go-api-template/internal/pkg/environment"
)

func main() {
	// Context for main function
	ctx := context.Background()
	ctx = customcontext.SetRequestID(ctx, "main-function")

	// Loading environment variables
	err := environment.LoadEnvironment()
	if err != nil {
		ctxlogger.Error(ctx, err.Error())
		return
	}

	// Setting chi router and serving it
	apiRouter := api.NewRouter()

	serverAddr := fmt.Sprintf(":%d", environment.PORT)
	ctxlogger.Info(ctx, "Starting server at %s", serverAddr)

	err = http.ListenAndServe(serverAddr, apiRouter)
	if err != nil {
		ctxlogger.Error(ctx, err.Error())
	}
}
