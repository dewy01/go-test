package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func message(ctx context.Context, ch chan int) {
	defer wg.Done()
	defer close(ch)

	for i := 0; i < 3; i++ {
		n := i
		select {
		case <-ctx.Done():
			return
		case ch <- n:
		}
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	quit := make(chan bool)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	wg.Add(2)
	go message(ctx, ch1)
	go message(ctx, ch2)
	go func() {
		wg.Wait()
		select {
		case <-ctx.Done():
		case quit <- true:
		}
	}()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("timeout")
			return
		case i, ok := <-ch1:
			if ok {
				fmt.Println("Received message from ch1: ", i)
			}
		case i, ok := <-ch2:
			if ok {
				fmt.Println("Received message from ch2: ", i)
			}
		case <-quit:
			fmt.Println("Quit invoked, jobs are done")
		}
	}
}
