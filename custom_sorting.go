package main

import (
	"fmt"
	"sort"
)

type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	b := []string{"me", "you", "I"}

	sort.Sort(ByLength(b))

	fmt.Println(b)

}
