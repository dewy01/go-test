package main

import (
	"fmt"
	"net/http"
	"time"
)

func fetchUrl(ch chan string, url string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("Error: %s", err)
		return
	}
	defer resp.Body.Close()
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("Response in %.2fs with code %d for %s", secs, resp.StatusCode, url)
}

func main() {
	ch := make(chan string)
	fmt.Println("start")
	urls := []string{"http://golang.org", "https://google.com", "https://youtube.com"}

	for _, url := range urls {
		go fetchUrl(ch, url)
	}

	for range urls {
		fmt.Println(<-ch)
	}
}
