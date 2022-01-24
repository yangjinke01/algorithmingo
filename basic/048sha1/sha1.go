package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	s := "s strings"
	hash := sha1.New()

	hash.Write([]byte(s))
	sum := hash.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x", sum)
}
