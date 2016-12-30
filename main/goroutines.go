package main

import "fmt"

func f(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s, "-", i)
	}
}

func main() {
	f("mdl")

	go f("gominsler")

	// for anonymous func
	go func(s string) {
		fmt.Println(s)
	}("dzm")

	var input string
	fmt.Scanln(&input)
	fmt.Println(input)
}
