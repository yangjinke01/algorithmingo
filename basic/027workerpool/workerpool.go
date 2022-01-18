package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println("work", id, "start", job)
		time.Sleep(time.Second)
		fmt.Println("work", id, "finish", job)
		results <- job * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 0; w < 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 0; j < numJobs; j++ {
		jobs <- j
	}
	close(jobs)
	//ensures that the worker goroutines have finished
	//alternative way to wait for multiple goroutines is to use a WaitGroup.
	for r := 0; r < numJobs; r++ {
		<-results
	}
}
