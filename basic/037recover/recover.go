package main

import "fmt"

func myPanic() {
	panic("a panic")
}

func main() {
	//defer will be active after a panic
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover from error:\n", err)
		}
	}()
	myPanic()
	//this code will not run
	fmt.Println("after myPanic")
}
