package main

import (
	"fmt"
	"time"
)

func demoUnbufferedChannel() {
	msgChan := make(chan string)

	go func() {
		time.Sleep(time.Duration(2) * time.Second)
		// receivers
		msgChan <- "hello"
		msgChan <- "world"
	}()

	// senders
	msg1 := <-msgChan
	msg2 := <-msgChan

	fmt.Println(msg1, msg2)
}
