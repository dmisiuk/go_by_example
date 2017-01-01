package main

import "fmt"

func zeroval(v int) {
	v = 3
}

func zeropt(v *int) {
	*v = 1
}

func main() {
	d := 10
	fmt.Println(d)

	zeroval(d)
	fmt.Println(d)

	zeropt(&d)
	fmt.Println(d)
}
