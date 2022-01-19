package main

import (
	"fmt"
	"sync"
)

type container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *container) inc(k string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[k]++
}

func main() {
	c := container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup
	doIncrement := func(k string, n int) {
		for i := 0; i < n; i++ {
			c.inc(k)
		}
		wg.Done()
	}
	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)
	wg.Wait()
	fmt.Println(c.counters)
}
