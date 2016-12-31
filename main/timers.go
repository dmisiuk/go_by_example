package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(time.Second)

	<-timer1.C
	fmt.Println("timer 1 expired")

	timer2 := time.NewTimer(time.Second * 2)

	messages := make(chan string)

	go func() {
		<-timer2.C
		fmt.Println("timer2 expired")
		messages <- "hello from anon"
	}()
	msg := <-messages
	fmt.Println(msg)
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("timer2 stopped")
	}
}
