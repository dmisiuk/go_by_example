package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() {
		messages <- "ping"
		messages <- "ping-2"
		msg := <-messages
		fmt.Println(msg)
	}()

	msg := <-messages
	fmt.Println(msg)

	msg = <-messages
	fmt.Println(msg)
	messages <- "from main thread"
}
