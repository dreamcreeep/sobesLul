package main

import (
	"fmt"
	"sync"
)

// Провести код-ревью и исправить ошибки.

func main() {
	var wg sync.WaitGroup
	occurrences := make(map[int]int)
	numbers := make(chan int, 100)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			for num := range numbers {
				occurrences[num]++
			}
		}()
	}

	for i := 1; i <= 100; i++ {
		numbers <- i
	}

	wg.Wait()
	close(numbers)

	fmt.Println("Occurrences:", occurrences)
}
