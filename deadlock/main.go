package main

import (
	"fmt"
	"strconv"
)

func toString(c chan string, val int) {
	if val > 5 {
		c <- strconv.Itoa(val)
	}
}

func main() {
	baseString := ""

	c1 := make(chan string)
	c2 := make(chan string)

	go toString(c1, 10)
	go toString(c2, 1)

	baseString += <-c1

	// Value in c2 doesnt pass toString requirement
	// fatal error: all goroutines are asleep - deadlock!
	baseString += <-c2

	fmt.Println(baseString)
}
