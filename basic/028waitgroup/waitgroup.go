package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int) {
	for job := range jobs {
		fmt.Println("work", id, "start", job)
		time.Sleep(time.Second)
		fmt.Println("work", id, "finish", job)
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	var wg sync.WaitGroup

	for w := 0; w < 3; w++ {
		wg.Add(1)
		//Avoid re-use of the same w value in each goroutine closure.
		w := w
		go func() {
			defer wg.Done()
			worker(w, jobs)
		}()
	}

	for j := 0; j < numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	wg.Wait()
}
