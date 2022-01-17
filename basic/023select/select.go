package main

import (
	"fmt"
	"time"
)

//Go’s select lets you wait on multiple channel operations.
//Combining goroutines and channels with select is a powerful feature of Go.
func main() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 2)

	go func() {
		time.Sleep(time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case one := <-c1:
			fmt.Println("received", one)
		case two := <-c2:
			fmt.Println("received", two)
		//Timeouts are important for programs that connect to external resources
		case <-time.After(time.Second * 1):
			fmt.Println("timeout 1 second")
		}
	}
}
