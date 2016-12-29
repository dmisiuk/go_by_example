package main

import "fmt"

func main() {
	var a [3]int
	fmt.Println("a:", a)

	a[0] = 2
	fmt.Println("after chaning:", a)

	fmt.Println("the second member", a[1])

	fmt.Println("len:", len(a))

	for i := 0; i < len(a); i++ {
		fmt.Println("a[",i,"]=", a[i])
	}
}
