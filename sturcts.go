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

func changecard(c *card) {
	c.cardUUID = "_"
	c.nfcUUID = "_"
}

func main() {
	c := card{"101", "sewqdfae1"}
	fmt.Println(c)
	printcard(c)
	fmt.Println(c)

	cp := &c
	changecard(cp)
	fmt.Println(cp)
	fmt.Println(c)

}
