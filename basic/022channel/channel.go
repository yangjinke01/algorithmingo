package main

import (
	"fmt"
	"time"
)

func work(done chan string) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- "done"
}

func main() {
	message := make(chan string)

	go func() { message <- "hello" }()
	msg := <-message
	fmt.Println(msg)

	//buffering up to 2 values
	messages := make(chan string, 2)

	messages <- "ping"
	messages <- "pong"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

	//synchronization
	done := make(chan string)
	go work(done)
	<-done
}
