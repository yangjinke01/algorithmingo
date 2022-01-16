package main

import "fmt"

func fact(num int) int {
	if num == 0 {
		return 1
	}
	return num * fact(num-1)
}

func main() {
	var fib func(num int) int

	fib = func(num int) int {
		if num < 2 {
			return num
		}
		return fib(num-2) + fib(num-1)
	}

	fmt.Println(fib(2))

	fmt.Println(fact(3))
}
