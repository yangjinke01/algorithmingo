package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	matched, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(matched)

	re, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println(re.MatchString("peach"))
	fmt.Println(re.FindString("peach punch"))
	fmt.Println(re.FindStringIndex("peach punch"))
	fmt.Println(re.FindStringSubmatch("peach punch"))
	fmt.Println(re.FindStringSubmatchIndex("peach punch"))
	fmt.Println(re.FindAllString("peach punch pinch", -1))
	fmt.Println(re.FindAllStringSubmatchIndex("peach punch pinch", -1))
	fmt.Println(re.FindAllString("peach punch pinch", 2))

	fmt.Println(re.Match([]byte("peach")))

	//MustCompile panics instead of returning an error
	re = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", re)

	fmt.Println(re.ReplaceAllString("a peach", "<fruit>"))

	upper := re.ReplaceAllFunc([]byte("peach"), bytes.ToUpper)
	fmt.Println(string(upper))
}
