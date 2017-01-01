package main

import "os"

func main() {
	//panic("a problem")

	_, err := os.Create("/bbbb/write")
	if err != nil {
		panic(err)
	}
}
