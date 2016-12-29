package main

import "fmt"

func main() {
	var s = make([]string,3)
	fmt.Println(s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "d"

	fmt.Println("after chaning:",s)

	fmt.Println("len:", len(s))

	s = append(s, "e")
	fmt.Println("after adding:", s)

	fmt.Println("slice 1:2", s[1:2])
}
