package main

import "fmt"

func tt() {
	a := []byte{'0', '1', '2'}
	b := append(a[1:2], 'a') // [1,a] len1, cap2
	b[0] = 'b'

	fmt.Printf("%s\n", a) // [0,b,a]
}

func main() {
	tt()
}
