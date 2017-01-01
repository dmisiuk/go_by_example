package main

import "fmt"

func main() {
	if b := 2; b < 3 {
		fmt.Println("b is less than 3, is", b)
	} else {
		fmt.Println("b is great or equeal to 3, is", b)
	}

	if 7%2 == 1 {
		fmt.Println("7 is not odd")
	}
}
