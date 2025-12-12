package main

import "fmt"

func a() {
	x := []int{}
	x = append(x, 0)  // [0] 1, 1
	x = append(x, 1)  // [0, 1] 2, 2
	x = append(x, 2)  // [0, 1, 2] 3, 4
	y := append(x, 3) // [0,1,2,3] 4, 4
	z := append(x, 4) // [0,1,2,4] 4, 4
	fmt.Println(y, z)
}

func main() {
	a() // какой здесь будет вывод?
}
