package main

import (
	"fmt"
	"strings"
)

var p = fmt.Println

func main() {
	h := "hello"
	p("index", strings.Index(h, "e"))
	p("split", strings.Split(h, ""))
	p("count", strings.Count(h, "l"))
	p("toLower", strings.ToLower(h))
	p("has prefix", strings.HasPrefix(h, "he"))
}
