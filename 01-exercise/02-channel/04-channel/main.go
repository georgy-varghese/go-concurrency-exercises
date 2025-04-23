package main

// TODO: Implement relaying of message with Channel Direction

func genMsg(wr chan<- int) {
	defer close(wr)
	// send message on ch1
	for i := 0; i < 10; i++ {
		wr <- i
	}
}

func relayMsg(rd <-chan int, wr chan<- int) {
	defer close(wr)
	// recv message from rd
	for i := range rd {
		// send it on wr
		wr <- i
	}

}

func main() {
	// create ch1 and ch2
	ch1 := make(chan int)
	ch2 := make(chan int)

	// spin goroutine genMsg and relayMsg
	go genMsg(ch1)

	// recv message on ch2
	go relayMsg(ch1, ch2)
	// print message on ch2
	for i := range ch2 {
		println(i)
	}
}
