package main

import "fmt"

func PPrintln(msg string, i *int) {
	fmt.Printf("%s %d\n", msg, *i)
}

func tt() (a int) {
	a = 1
	defer func() {
		a++
	}()

	defer func() {
		a = 5
	}()

	return 3
}

func main() {
	a := 1
	b := 2
	pb := &b

	defer fmt.Println("a=", a)
	defer fmt.Println("b=", *pb)
	defer PPrintln("pb=", pb)

	a = 10
	*pb = 20

	fmt.Println(tt())
}

// 6
// pb=20
// b=2
// a=1
