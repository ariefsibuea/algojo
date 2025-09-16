package main

import (
	"fmt"
	"time"
)

func demoBufferedChannel() {
	msgChan := make(chan string, 3)

	go func() {
		time.Sleep(time.Duration(2) * time.Second)
		// sender
		msgChan <- "Arief"
		msgChan <- "Edy"
		msgChan <- "Putra"
	}()

	// receiver
	msg1 := <-msgChan
	msg2 := <-msgChan
	msg3 := <-msgChan

	fmt.Println(msg1, msg2, msg3)
}
