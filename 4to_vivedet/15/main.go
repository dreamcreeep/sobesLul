package main

import (
	"fmt"
)

func test() (x int) {
	x = 5
	defer func() {
		x = x + 10
	}()

	return x
}

func main() {
	fmt.Println(test())
}
