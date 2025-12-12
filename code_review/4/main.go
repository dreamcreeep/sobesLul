package main

import (
	"fmt"
	"strconv"
	"strings"
)

func fn() {
	var str strings.Builder

	for i := 0; i < 10000; i++ {
		str. /*.AddString*/ WriteString(strconv.Itoa(i))
	}

	fmt.Println(str)
}

func main() {
	fn()
}
