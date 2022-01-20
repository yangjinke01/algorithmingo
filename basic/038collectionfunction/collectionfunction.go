package main

import (
	"fmt"
	"strings"
)

func Index(items []string, target string) int {
	for index, value := range items {
		if value == target {
			return index
		}
	}
	return -1
}

func Include(items []string, target string) bool {
	return Index(items, target) >= 0
}

func Any(items []string, f func(string) bool) bool {
	for _, item := range items {
		if f(item) {
			return true
		}
	}
	return false
}

func All(items []string, f func(string) bool) bool {
	for _, item := range items {
		if !f(item) {
			return false
		}
	}
	return true
}

func Filter(items []string, f func(string) bool) []string {
	var result []string
	for _, item := range items {
		if f(item) {
			result = append(result, item)
		}
	}
	return result
}

func Map(items []string, f func(string) string) []string {
	var result []string
	for _, item := range items {
		result = append(result, f(item))
	}
	return result
}

func main() {
	fruits := []string{"peach", "apple", "pear", "plum"}

	fmt.Println(Index(fruits, "apple"))

	fmt.Println(Include(fruits, "pear"))

	fmt.Println(Any(fruits, func(fruit string) bool {
		return strings.HasPrefix(fruit, "a")
	}))

	fmt.Println(All(fruits, func(fruit string) bool {
		return strings.HasPrefix(fruit, "p")
	}))

	fmt.Println(Filter(fruits, func(fruit string) bool {
		return strings.HasPrefix(fruit, "p")
	}))

	fmt.Println(Map(fruits, func(fruit string) string {
		return strings.ToUpper(fruit)
	}))
}
