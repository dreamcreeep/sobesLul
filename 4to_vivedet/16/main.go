package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["a"] = 1

	p := &m["a"]
	*p = 42

	fmt.Println(m)
}
