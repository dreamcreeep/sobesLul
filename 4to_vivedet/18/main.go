package main

import (
	"fmt"
	"sync"
)

// 4то выведется и предложить исправление
func main() {
	c := make(chan int)

	wg := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(c chan int) {
			defer wg.Done()

			foo(c)
		}(c)
	}

	sum := 0
	for r := range c {
		sum += r
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	fmt.Println(sum)

}

func foo(c chan int) {

	for i := 0; i < 5; i++ {
		c <- i
	}
}
