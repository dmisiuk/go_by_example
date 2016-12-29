package main

import "fmt"

func main() {
	var m map[string]string = make(map[string]string)

	m["s"] = "d"
	m["1"] = "first"
	fmt.Println(m)

	delete(m, "1")

	fmt.Println("after delete:", m)

	k, v := m["s2"]

	fmt.Println("k:", k)
	fmt.Println("v:", v)

	inline_m := map[string]int{"1":1, "2":2}

	fmt.Println("inline_m:", inline_m)
}
