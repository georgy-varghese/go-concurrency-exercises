package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// what is the output
	//TODO: fix the issue.

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		// Passing i as variable to goroutine
		// In tutorial, output without this change is 4 4 4
		// It might seem there is no difference in output with/without this change
		// However, not passing i as argument to goroutine, program is error prone and unreliable
		// On passing the argument, a copy of i is created for each goroutine
		// so value is not affected by loops subsequent iteration
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
}
