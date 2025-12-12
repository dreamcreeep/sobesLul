package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	_, _ = worker(), worker()

	fmt.Printf("time passed %v", time.Since(start))
}

func worker() <-chan int {
	ch := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- 1
	}()

	return ch
}
