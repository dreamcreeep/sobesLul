// Что выведет программа?

package main

import "fmt"

func main() {
	s := "hêllo"
	for i := range s {
		fmt.Printf("position %d: %c\n", i, s[i])
	}
	fmt.Printf("len=%d\n", len(s))
}

// h:0,E:1,l:3,l:4,o:5 len=6
