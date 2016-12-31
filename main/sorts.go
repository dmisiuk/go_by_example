package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"b", "e", "ab", "ac"}
	sort.Strings(s)
	fmt.Println("sorted:", s)

	ints := []int{1, 3, 0, -5}

	sort.Ints(ints)
	fmt.Println("ints:", ints)

	isSotred := sort.IntsAreSorted(ints)
	fmt.Println("isSorted:", isSotred)
}
