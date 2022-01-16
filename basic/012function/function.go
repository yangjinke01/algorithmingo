package main

import "fmt"

func plus(a, b int) int {
	return a + b
}

func plusplus(a, b, c int) int {
	return a + b + c
}

func values() (int, int) {
	return 3, 4
}
func sum(nums ...int) int {
	result := 0
	for _, num := range nums {
		result += num
	}
	return result
}
func main() {
	fmt.Println("1+2 =", plus(1, 2))
	fmt.Println("1+2+3 =", plusplus(1, 2, 3))

	v1, v2 := values()
	fmt.Println("values:", v1, v2)

	fmt.Println("sum:", sum(2, 3, 4))
}
