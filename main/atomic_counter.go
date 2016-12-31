package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var counter uint64 = 0

	for i := 0; i < 50; i++ {
		go func() {
			for {
				atomic.AddUint64(&counter, 1)
				time.Sleep(time.Nanosecond)
			}
		}()
	}
	time.Sleep(time.Second)

	value := atomic.LoadUint64(&counter)
	fmt.Println(value)
}
