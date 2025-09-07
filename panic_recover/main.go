package main

import (
	"fmt"
	"sync"
)

type listSKU []string

func (l listSKU) getLastSKU() string {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recovered from panic: %v", r)
		}
	}()

	return l[len(l)]
}

func main() {

	items := listSKU{
		"MP990099991",
		"MP900000002",
		"MP000000003",
		"MP000000004",
		"MP000000005",
	}
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		lastItem := items.getLastSKU()
		fmt.Printf("Last SKU is: %s\n", lastItem)
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("Program completed.")
}
