package main

import "fmt"

type Seller interface {
	GetId() int64
}

type Lamoda struct{}

func (l Lamoda) GetId() int64 {
	return 1
}

func main() {
	var seller Seller

	if seller == nil {
		fmt.Println("nil interface")
	}

	var lamoda *Lamoda

	if lamoda == nil {
		fmt.Println("nil struct")
	}

	seller = lamoda

	if seller == nil {
		fmt.Println("vas naebali")
	}
}
