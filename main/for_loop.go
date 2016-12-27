package main

import "fmt"

func main() {
	i := 0
	for {
		fmt.Println(i)
		i++
		break
	}

	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for j := 100; j > 0; j-- {
		fmt.Print(j, " ")
	}
}
