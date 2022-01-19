package main

import "fmt"

func myPanic() {
	panic("a panic")
}

func main() {
	//defer will be active after a panic
	defer func() {
		fmt.Println("defer")
	}()
	myPanic()
	fmt.Println("after myPanic")
}
