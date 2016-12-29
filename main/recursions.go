package main

import "fmt"

func fact(n int) int {
	if n == 1 {
		return 1
	}

	return n * fact(n - 1)
}

func main() {
	fmt.Println("1!=", fact(1))
	fmt.Println("5!=", fact(5))
}
