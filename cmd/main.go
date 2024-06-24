package main

import (
	"net/http"

	"github.com/Suhaan-Bhandary/go-api-template/internal/api"
)

func main() {
	apiRouter := api.NewRouter()
	http.ListenAndServe(":8080", apiRouter)
}
