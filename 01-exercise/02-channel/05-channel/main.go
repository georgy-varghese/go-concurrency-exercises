package main

import "fmt"

func main() {
	// Create channel owner goroutine which return channel and
	// writes data into channel and
	// closes the channel when done.
	owner := func() <-chan int {
		ch := make(chan int)
		go func() {
			defer close(ch) // close channel when done
			// write values to channel
			for i := 0; i < 5; i++ {
				ch <- i
			}
		}()
		fmt.Println("Done sending!")
		return ch
	}

	consumer := func(ch <-chan int) {
		// read values from channel
		for v := range ch {
			fmt.Printf("Received: %d\n", v)
		}
		fmt.Println("Done receiving!")
	}

	ch := owner()
	consumer(ch)
}
