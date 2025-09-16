package main

import (
	"fmt"
	"sync"
	"time"
)

func demoFanOutFanIn() {
	input := make(chan int, 10)

	// fan-out: distribute work to multiple workers
	worker1 := worker(input, 1)
	worker2 := worker(input, 2)
	worker3 := worker(input, 3)

	// fan-in: merge results from multiple workers
	output := fanIn(worker1, worker2, worker3)

	// send work
	go func() {
		for i := 0; i < 10; i++ {
			input <- i
		}
		close(input)
	}()

	// collect results
	for result := range output {
		fmt.Printf("%v\n", result)
	}
}

func worker(input <-chan int, workerID int) <-chan string {
	output := make(chan string)

	go func() {
		defer close(output)

		for n := range input {
			time.Sleep(time.Second) // simulate work
			output <- fmt.Sprintf("Worker %d receive input: %d, task result: %d", workerID, n, n*n)
		}
	}()

	return output
}

func fanIn(inputs ...<-chan string) <-chan string {
	output := make(chan string)
	wg := sync.WaitGroup{}

	for _, input := range inputs {
		wg.Add(1)

		go func(ch <-chan string) {
			defer wg.Done()

			for value := range ch {
				output <- value
			}
		}(input)
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}
