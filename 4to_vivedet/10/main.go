package main

import "fmt"

// Что выведет?
func main() {

	d := make(map[string]struct{})
	a := []string{
		"one", "two", "three", "foo", "bar",
	}

	for _, s := range a {
		d[s] = struct{}{}
	}

	for k := range d {
		fmt.Println(k)
	}
}

// [one:{}],[two:{}]
