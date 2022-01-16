package main

import "fmt"

func zeroValue(num int) {
	num = 0
}
func zeroPointer(num *int) {
	*num = 0
}
func main() {
	num := 1
	zeroValue(num)
	fmt.Println("zeroValue", num)

	zeroPointer(&num)
	fmt.Println("zeroPointer", num)

	fmt.Println("pointer:", &num)
}
