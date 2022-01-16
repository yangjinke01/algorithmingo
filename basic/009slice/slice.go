package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("append", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)

	l := s[2:5]
	fmt.Println("slice1", l)

	l = s[:5]
	fmt.Println("slice2", l)

	l = s[2:]
	fmt.Println("slice3", l)

	twoD := make([][]int, 3)
	for i := 0; i < len(twoD); i++ {
		twoD[i] = make([]int, i+1)
		for j := 0; j < len(twoD[i]); j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("twoD:", twoD)
}
