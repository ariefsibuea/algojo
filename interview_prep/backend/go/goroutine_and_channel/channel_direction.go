package main

import "fmt"

// send-only channel parameter
func producer(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
}

// receive-only channel parameter
func consumer(ch <-chan int) {
	for value := range ch {
		fmt.Printf("Consumed: %d\n", value)
	}
}

func demoChannelDirection() {
	ch := make(chan int, 2)

	go producer(ch)
	consumer(ch)
}
