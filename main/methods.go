package main

import "fmt"

type rect struct {
	width, height int
}

func (r rect) perimeter() int {
	return 2 * (r.height + r.width)
}

func (r *rect) area() int {
	return r.height * r.width
}

func main() {
	r := rect{width:20, height:5}
	fmt.Println(r)

	fmt.Println("area:", r.area())
	fmt.Println("perimeter:", r.perimeter())

}
