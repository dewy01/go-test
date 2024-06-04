package api

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"golang.org/x/net/websocket"
)

func StartHTTPServer(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	mux := http.NewServeMux()
	s := NewServer()

	mux.Handle("/ws", websocket.Handler(s.HandleWS))

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		fmt.Println("Starting HTTP Server...")
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			fmt.Printf("HTTP Server error: %s \n", err)
		}
	}()

	<-ctx.Done()
	fmt.Println("Shutting down HTTP Server...")
	shutdownCtx, cancelShutdwn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdwn()

	err := server.Shutdown(shutdownCtx)
	if err != nil {
		fmt.Printf("HTTP Server shtdown error: %s \n", err)
	}

	fmt.Println("HTTP Server stopped")
}
