package main

import (
	"context"
	"fmt"
	"hero/api"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Add(1)
	go api.StartHTTPServer(ctx, &wg)

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	<-signalCh

	fmt.Println("Shutting down HTTP server gracefully...")
	cancel()

	wg.Wait()

	fmt.Println("Shutdown completed.")
}
