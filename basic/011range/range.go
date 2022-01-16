package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	for _, num := range nums {
		fmt.Println(num)
	}

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	for k := range kvs {
		fmt.Println("k:", k)
	}
	for i, c := range "jack" {
		fmt.Println(i, c)
	}
}
