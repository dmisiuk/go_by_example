package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received jobs", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for i := 0; i < 3; i++ {
		fmt.Println("send jobs", i)
		jobs <- i
	}
	close(jobs)
	fmt.Println("sent all jobs")
	<-done
}
