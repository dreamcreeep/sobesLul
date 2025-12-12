package main

// Необходимо написать программу, в которой N горутин одновременно пишут в канал любое число,
// а main-горутина находит сумму всех чисел, записанных в канал
import (
	"fmt"
	"math/rand"
	"sync"
)

func foo(n int) <-chan int {
	outputChan := make(chan int, n)

	wg := &sync.WaitGroup{}

	for range n {
		wg.Add(1)
		go func() {
			defer wg.Done()
			num := rand.Int()
			outputChan <- num
		}()
	}

	go func() {
		wg.Wait()
		close(outputChan)
	}()

	return outputChan
}

func main() {
	sumChan := foo(5)

	sum := 0

	for num := range sumChan {
		sum += num
	}

	fmt.Printf("digits summary: %v", sum)
}
