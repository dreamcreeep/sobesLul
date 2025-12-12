package main

import (
	"log"
)

type MyStruct struct{}

func (s MyStruct) Hello() {
	log.Println("Hi")
}

func main() {
	a := MyStruct{}

	b := &a

	b.Hello()

}
