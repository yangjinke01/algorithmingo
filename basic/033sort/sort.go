package main

import (
	"fmt"
	"sort"
)

func main() {
	strings := []string{"a", "c", "b", "d"}
	sort.Strings(strings)
	fmt.Println("Strings:", strings)

	nums := []int{1, 3, 2, 5}
	sort.Ints(nums)
	fmt.Println("Nums:", nums)

	s := sort.IntsAreSorted(nums)
	fmt.Println("Sorted:", s)
}
