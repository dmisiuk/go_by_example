package main

import (
	"fmt"
	"os"
)

func main() {
	f := createFile("/tmp/my_go_file.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(name string) *os.File {
	fmt.Println("creating file")
	f, err := os.Create(name)
	if err != nil {
		panic(err)
	}

	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing file")
	fmt.Fprintln(f, "some data")
}

func closeFile(f *os.File) {
	fmt.Println("closing file")
	f.Close()
}
