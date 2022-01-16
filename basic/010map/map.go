package main

import "fmt"

func main() {
	m := make(map[string]int, 2)
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("m:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)
	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("len:", len(m))
	fmt.Println("m:", m)

	_, present := m["k1"]
	fmt.Println("present:", present)

	n := map[string]int{"name": 1, "sex": 2}
	fmt.Println("map:", n)
}
