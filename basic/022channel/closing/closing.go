package main

import (
	"fmt"
)

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			//the more value will be false if jobs has been closed and
			//all values in the channel have already been received
			job, more := <-jobs
			if more {
				fmt.Println("received job,", job)
			} else {
				fmt.Println("received all job")
				done <- true
				break
			}
		}
	}()

	for i := 1; i <= 3; i++ {
		jobs <- i
		fmt.Println("send job,", i)
	}
	close(jobs)
	fmt.Println("send all jobs")
	<-done
}
