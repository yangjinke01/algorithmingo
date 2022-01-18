package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan string)

	select {
	case msg := <-messages:
		fmt.Println("received message,", msg)
	default:
		fmt.Println("no message received")
	}

	//Here msg cannot be sent to the messages channel,
	//because the channel has no buffer and there is no receiver.
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("send message,", msg)
	default:
		fmt.Println("no message send")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message,", msg)
	case sig := <-signals:
		fmt.Println("received signals,", sig)
	default:
		fmt.Println("no activity")
	}
}
