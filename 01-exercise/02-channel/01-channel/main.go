package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go func(a, b int) {
		defer close(ch)
		c := a + b
		ch <- c
	}(1, 2)
	// get the value computed from goroutine
	r := <-ch
	// get the value from channel
	fmt.Printf("computed value %v\n", r)
}
