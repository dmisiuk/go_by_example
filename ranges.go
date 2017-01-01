package main

import "fmt"

func main() {
	a := []int{1, 2, 3}

	fmt.Println(a)

	sum := 0

	for _, num := range a {
		sum += num
	}

	fmt.Println("sum:", sum)

	m := map[string]string{"1": "first", "2": "second"}

	for k, v := range m {
		fmt.Println(k, "->", v)
	}

	for i, c := range "умный" {
		fmt.Println(i, "->", c)
	}

}
