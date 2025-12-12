package main

import "fmt"

func main() {
	a := make([]int, 0, 4) // [0,0,0,0] 0,4
	a = append(a, 1, 2, 3) // [1,2,3,0] 3,4
	appendSlice(a)
	fmt.Println(a[:]) // [1,2,3]
	fmt.Println(len(a), cap(a))
	fmt.Println(a[:4])
}

func appendSlice(a []int) {
	a = append(a, 4)

	fmt.Println(len(a), cap(a))
}
