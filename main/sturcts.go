package main

import "fmt"

type card struct {
	cardUUID string
	nfcUUID  string
}

func printcard(c card) {
	fmt.Println("cardUUID:", c.cardUUID)
	fmt.Println("nfcUUID:", c.nfcUUID)
	c.cardUUID = "2"
}

func main() {
	c := card{"101", "sewqdfae1"}
	fmt.Println(c.cardUUID)
	printcard(c)
	fmt.Println("after invoking:", c.cardUUID)
}
