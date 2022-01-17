package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go f("goroutine")

	//You can also start a goroutine for an anonymous function call
	go func(message string) {
		fmt.Println(message)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
