package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID   int
	Data string
}

type Result struct {
	JobID int
	Value string
	Error error
}

func workerPool(numWorkers int, jobs <-chan Job, results chan<- Result) {
	var wg sync.WaitGroup

	// start workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			// loop over receiver
			for job := range jobs {
				time.Sleep(time.Second) // simulate work
				results <- Result{
					JobID: job.ID,
					Value: fmt.Sprintf("Worker %d processed job %d", workerID, job.ID),
					Error: nil,
				}
			}
		}(i)
	}

	go func() {
		wg.Wait()
		// close results to prevent goroutines asleep and cause a deadlock
		close(results)
	}()
}

func demoWorkerPool() {
	jobs := make(chan Job, 4)
	results := make(chan Result, 4)

	workerPool(2, jobs, results) // all workers share the same job channel: 'jobs'

	go func() {
		for i := 0; i < 20; i++ {
			jobs <- Job{
				ID:   i,
				Data: fmt.Sprintf("Send job %d", i),
			}
		}
		// close job to prevent goroutines asleep and cause a deadlock
		close(jobs)
	}()

	for result := range results {
		if result.Error != nil {
			fmt.Printf("Found error = %v\n", result.Error)
			continue
		}

		fmt.Printf("%v\n", result.Value)
	}
}
