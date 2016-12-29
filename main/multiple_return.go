package main

import "fmt"

func main() {
	m, n := returnPair(1, 2)
	fmt.Println("m:", m, "n:", n)

	_, s := returnPair(10, 23)
	fmt.Println("s:", s)

}

func returnPair(a, b int) (int, int) {
	return a, b
}
