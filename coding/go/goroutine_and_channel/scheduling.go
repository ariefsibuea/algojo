package main

import (
	"fmt"
	"runtime"
	"time"
)

func demoScheduling() {
	runtime.GOMAXPROCS(2) // limit to 2 OS threads

	for i := 0; i < 10; i++ {
		go func(id int) {
			// CPU-bound work that will cause scheduling
			for j := 0; j < 1000000; j++ {
				_ = j * j
			}
			fmt.Printf("Goroutine %d completed on thread %d\n", id, runtime.NumGoroutine())
		}(i)
	}

	time.Sleep(time.Second)
}
