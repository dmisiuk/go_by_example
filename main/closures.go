package main

import "fmt"

func main() {
	ni := nextInt()
	fmt.Println(ni())
	fmt.Println(ni())
	fmt.Println(ni())

	nni := nextInt()
	fmt.Println(nni())
}

func nextInt() func() int {
	i := 0
	return func() int {
		i++
		return i;
	}
}
