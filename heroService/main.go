package main

import (
	"hero/api"
	"hero/middleware"
	"net/http"
)

func main() {
	mux := api.NewApp()

	stack := middleware.CreateStack(
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8080",
		Handler: stack(mux),
	}

	server.ListenAndServe()
}
