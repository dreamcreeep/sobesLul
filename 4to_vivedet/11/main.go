package main

import "fmt"

// Что выведет?
func main() {
	x := 1
	defer fmt.Println("first defer", x)
	x = 2
	fmt.Println("value", x)
	x = 3
	defer fmt.Println("second defer", x)
	x = 4
}
