package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)
	for request := range requests {
		<-limiter
		fmt.Println("request:", request, time.Now())
	}

	burstLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstLimiter <- time.Now()
	}
	go func() {
		for tick := range time.Tick(200 * time.Millisecond) {
			burstLimiter <- tick
		}
	}()

	burstRequests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		burstRequests <- i
	}
	close(burstRequests)

	for request := range burstRequests {
		<-burstLimiter
		fmt.Println("request:", request, time.Now())
	}
}
