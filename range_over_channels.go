package main

import "fmt"

func main() {
	queue := make(chan string, 2)

	queue <- "first"
	queue <- "second"
	close(queue)

	for m := range queue {
		fmt.Println(m)
	}

}
