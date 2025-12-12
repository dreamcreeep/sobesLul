// Что выведет программа?

package main

import (
	"fmt"
)

func main() {
	a := 10

	defer func() {
		fmt.Println("call 0 ", a+10)
	}()

	defer fmt.Println("call 1 ", a+10)

	a++

	fmt.Println("call 2 ", a)
}
