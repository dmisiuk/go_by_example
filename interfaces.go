package main

import "fmt"

type geometry interface {
	perimeter() float64
	area() float64
}

type rect struct {
	width, height float64
}

func (r *rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (r *rect) area() float64 {
	return r.width * r.height
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("perimeter:", g.perimeter())
	fmt.Println("area:", g.area())
}

func main() {
	r := rect{width: 10, height: 4}
	measure(&r)
}
