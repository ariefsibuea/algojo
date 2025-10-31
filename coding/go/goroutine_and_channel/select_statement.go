package main

import (
	"context"
	"fmt"
	"time"
)

func demoSelectPatterns(ctx context.Context) {
	ch1 := make(chan string)
	ch2 := make(chan string)
	timeout := time.After(5 * time.Second)

	go func() {
		time.Sleep(2 * time.Second) // simulate work
		ch1 <- "work 1 completed"
	}()

	go func() {
		time.Sleep(3 * time.Second) // simulate work
		ch2 <- "work 2 completed"
	}()

	// wait for result, timeout, or context cancellation
	// i := 1
	for {
		// fmt.Printf("Loop number-%d\n", i)

		select {
		case msg1 := <-ch1:
			fmt.Printf("Received: %s\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("Received: %s\n", msg2)
		case <-timeout:
			fmt.Printf("Operation timed out\n")
			return
		case <-ctx.Done():
			fmt.Printf("Context cancelled\n")
			return
		default:
			// The default case makes select non-blocking, but use it carefully as it can create busy-waiting loops.
			fmt.Printf("No channels ready, doing other work...\n")
			time.Sleep(500 * time.Millisecond)
		}

		// i++
	}
}
