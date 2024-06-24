package main

import (
	"fmt"
	"net/http"

	"github.com/Suhaan-Bhandary/go-api-template/internal/api"
)

func main() {
	apiRouter := api.NewRouter()

	err := http.ListenAndServe(":8080", apiRouter)
	if err != nil {
		fmt.Println(err)
	}
}
