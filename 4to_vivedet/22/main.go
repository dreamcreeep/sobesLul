package main

import "fmt"

func testSlices1() {
	a := []byte{'a', 'b', 'c'}
	// ['b'] len1, cap2
	b := append(a[1:2], 'd')
	// ['b','d'] len2, cap2
	b[0] = 'z'
	// ['z','d'] len2, cap2
	fmt.Printf("%s\n", a)
}

func main() {
	testSlices1()
}
