package main

import (
	"fmt"
	"os"
)

// Что выведет?

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	if err := Foo(); err != nil {
		fmt.Println("error")
	} else {
		fmt.Println("ok")
	}
}
