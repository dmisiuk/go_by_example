package main

import (
	"time"
	"fmt"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "c1"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "c2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-c1:
			fmt.Println(msg)
		case msg := <-c2:
			fmt.Println(msg)
		}
	}

}
