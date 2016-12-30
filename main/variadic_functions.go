package main

import "fmt"

func main() {
	fmt.Println("sum(1,2)=", sum(1, 2))
	fmt.Println("sum(1,2,3)=", sum(1, 2, 3))

	s := []int{10, 20, 40}
	fmt.Println("sum of ", s, "=", sum(s...))

}

func sum(nums ...int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}
