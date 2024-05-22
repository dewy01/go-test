package api

import (
	"context"
	"fmt"
	"hero/hero"
	"hero/middleware"
	"net/http"
	"sync"
	"time"
)

func StartHTTPServer(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	mux := http.NewServeMux()
	store := hero.NewHeroServer()

	mux.HandleFunc("POST /hero/add", store.CreateHeroHandler)
	mux.HandleFunc("PATCH /hero/update/{id}", store.UpdateHeroHandler)
	mux.HandleFunc("GET /hero/getAll", store.GetHeroesHandler)
	mux.HandleFunc("GET /hero/get/{id}", store.GetHeroByIdHandler)
	mux.HandleFunc("GET /hero/winner/{id}/{id2}", store.GetWinnerHandler)
	mux.HandleFunc("GET /hero/winner/all", store.GetGloblaWinnerHandler)
	mux.HandleFunc("DELETE /hero/delete/{id}", store.DeleteHeroHandler)

	stack := middleware.CreateStack(
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8080",
		Handler: stack(mux),
	}

	go func() {
		fmt.Println("Starting HTTP server...")
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			fmt.Printf("HTTP Server error: %s\n", err)
		}
	}()

	<-ctx.Done()
	fmt.Println("Shutting down HTPP server...")
	shutdownCtx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()

	err := server.Shutdown(shutdownCtx)
	if err != nil {
		fmt.Printf("HTTP Server shutdown error: %s\n", err)
	}

	fmt.Println("HTTP Server stopped")
}
