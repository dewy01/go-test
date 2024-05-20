package main

import (
	"fmt"
	"sync"
	"time"
)

func printHello() {
	fmt.Println("Hello")
}

func printMultiple(values []string) {
	for i := 0; i <= len(values)-1; i++ {
		fmt.Println(values[i])
	}
}
func inputSting(wg *sync.WaitGroup, c chan string, value string) {
	wg.Done()
	c <- value
}

func printString(wg *sync.WaitGroup, c chan string) {
	wg.Done()
	fmt.Println(len(c))
	for i := 0; i < len(c); i++ {
		fmt.Println(<-c)
	}
}

func main() {
	var wg sync.WaitGroup
	c := make(chan string, 3)
	wg.Add(1)
	go inputSting(&wg, c, "Message1")
	wg.Add(1)
	go inputSting(&wg, c, "Message2")
	wg.Add(1)
	go inputSting(&wg, c, "Message3")
	wg.Add(1)
	go printString(&wg, c)
	//go printHello()
	//go printMultiple([]string{"First", "Second", "Third"})
	time.Sleep(2000)
	wg.Wait()
	fmt.Println("World")

}
