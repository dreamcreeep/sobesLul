package main

import "fmt"

// Что выведет?

func foo(a []int) {
	//a = append(a, 5)

	for i := range a {
		a[i] = 20
	}
}

func main() {

	sl := []int{1, 2, 3, 4}
	foo(sl)

	fmt.Println(sl)

}
