package main

import (
	"fmt"
	"time"
)

func worker(i int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("start working with", j, "in job", i)
		time.Sleep(time.Second)
		fmt.Println("ended working with", j, "in job", i)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	for j := 0; j < 3; j++ {
		go worker(j, jobs, results)
	}

	for job := 0; job < 10; job++ {
		jobs <- job
	}

	for r := 0; r < 10; r++ {
		result := <-results
		fmt.Println("result:", result)
	}
}
