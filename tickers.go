package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Millisecond * 200)

	timer := time.NewTimer(time.Second * 2)

	go func() {
		for range ticker.C {
			fmt.Println("Ticker")
		}
	}()

	<-timer.C
	ticker.Stop()
	fmt.Println("ticker stopped")
}
