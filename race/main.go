package main

import (
	"fmt"
	"sync"
)

var (
	wg     sync.WaitGroup
	mu     sync.Mutex
	global string = ""
)

// To get rid of race error, add mutex Lock

func append(value string, isLocked bool) {
	switch isLocked {
	case false:
		global += value
	case true:
		mu.Lock()
		global += value
		mu.Unlock()
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go append("string1", false)
	go append("string2", false)

	wg.Wait()
	fmt.Println(global)
}
