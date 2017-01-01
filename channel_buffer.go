package main

import "fmt"

func main() {
	messages := make(chan string, 2)

	messages <- "ping-1"
	messages <- "ping-2"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
