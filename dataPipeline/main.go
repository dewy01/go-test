package main

import (
	"fmt"
	"strings"
	"sync"
)

func toUpperCase(input <-chan string, output chan<- string) {
	for str := range input {
		output <- strings.ToUpper(str)
	}
	close(output)
}

func addExclamation(input <-chan string, output chan<- string) {
	for str := range input {
		output <- str + "!"
	}
	close(output)
}

func main() {
	words := []string{"hello", "world", "golang", "concurrency"}

	ch1 := make(chan string, len(words))
	ch2 := make(chan string, len(words))
	ch3 := make(chan string, len(words))

	go func() {
		for _, word := range words {
			ch1 <- word
		}
	}()

	go toUpperCase(ch1, ch2)
	go addExclamation(ch2, ch3)

	var wg sync.WaitGroup
	wg.Add(len(words))

	go func() {
		for word := range ch3 {
			fmt.Println(word)
			wg.Done()
		}
	}()

	wg.Wait()
}
