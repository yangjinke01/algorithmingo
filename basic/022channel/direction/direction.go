package main

import "fmt"

//When using channels as function parameters,
//you can specify if a channel is meant to only send or receive values.
//This specificity increases the type-safety of the program.
func ping(pings chan<- string, message string) {
	pings <- message
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
