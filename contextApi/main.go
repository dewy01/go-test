package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func fetchAPI(ctx context.Context, url string, results chan<- string) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		results <- fmt.Sprintf("Error creating request for %s: %s", url, err.Error())
		return
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		results <- fmt.Sprintf("Error making request to %s: %s", url, err.Error())
		return
	}
	defer resp.Body.Close()

	results <- fmt.Sprintf("Response from %s: %d", url, resp.StatusCode)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	urls := []string{
		"https://catfact.ninja/fact",
		"https://www.boredapi.com/api/activity",
	}

	results := make(chan string)

	for _, url := range urls {
		go fetchAPI(ctx, url, results)
	}

	for range urls {
		fmt.Println(<-results)
	}
}
