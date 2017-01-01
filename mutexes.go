package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var readOps int32 = 0
	var writeOps int32 = 0
	m := make(map[int]int)

	mutex := sync.Mutex{}
	total := 0

	for i := 0; i < 10; i++ {
		go func() {
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += m[key]
				mutex.Unlock()
				atomic.AddInt32(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for i := 0; i < 100; i++ {
		go func() {
			for {
				key := rand.Intn(5)
				value := rand.Intn(100)
				mutex.Lock()
				m[key] = value
				mutex.Unlock()
				atomic.AddInt32(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	fmt.Println("read ops:", atomic.LoadInt32(&readOps))
	fmt.Println("write ops:", atomic.LoadInt32(&writeOps))
	fmt.Println("total:", total)
	fmt.Println("map", m)
}
