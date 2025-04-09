package main

import (
	"fmt"
	"sync"
	"time"
)

func fun(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	// Using wg wait
	var wg sync.WaitGroup
	wg.Add(1)

	// Direct call
	fun("direct call")

	// TODO: write goroutine with different variants for function call.

	// goroutine function call
	go fun("1 go routine")

	// goroutine with anonymous function
	go func() {
		defer wg.Done()
		fun("2 go routine")
	}()

	// goroutine with function value call
	val := fun
	val("3 go routine")

	// wait for wg
	wg.Wait()

	fmt.Println("done..")
}
