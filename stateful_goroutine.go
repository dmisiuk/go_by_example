package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key   int
	value int
	resp  chan bool
}

func main() {
	reads := make(chan readOp)
	writes := make(chan writeOp)
	var readOps int32 = 0
	var writeOps int32 = 0

	go func() {
		state := make(map[int]int)
		for {
			select {
			case read := <-reads:
				value := state[read.key]
				read.resp <- value
			case write := <-writes:
				state[write.key] = write.value
				write.resp <- true
			}
		}
	}()

	for i := 0; i < 100; i++ {
		go func() {
			for {
				wOp := writeOp{key: i, value: rand.Intn(1000), resp: make(chan bool)}
				writes <- wOp
				<-wOp.resp
				atomic.AddInt32(&writeOps, 1)
			}
		}()
	}

	for i := 0; i < 100; i++ {
		go func() {
			for {
				rOp := readOp{key: i, resp: make(chan int)}
				reads <- rOp
				<-rOp.resp
				atomic.AddInt32(&readOps, 1)
			}
		}()
	}

	time.Sleep(time.Second)

	fmt.Println("reads:", atomic.LoadInt32(&readOps))
	fmt.Println("writes:", atomic.LoadInt32(&writeOps))

}
