package main

import "fmt"

func ping(pings chan <- string, msg string) {
	fmt.Println("send message to pings")
	pings <- msg
}

func pong(pings <-chan string, pongs chan <- string) {
	fmt.Println("getting msg from pings")
	msg := <-pings
	fmt.Println("sending msg to pongs")
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "hi")
	pong(pings, pongs)

	fmt.Println(<-pongs)
}
