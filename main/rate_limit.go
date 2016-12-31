package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)

	for i := 0; i < 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.NewTicker(time.Millisecond * 300)

	for req := range requests {
		<-limiter.C
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.NewTicker(time.Millisecond * 500).C {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for r := range burstyRequests {
		limiterTime := <-burstyLimiter
		fmt.Println("request", r, "limiter time:", limiterTime, "now:", time.Now())
	}
}
