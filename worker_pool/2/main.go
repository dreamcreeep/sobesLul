package main

import (
	"fmt"
	"sync"
	"time"
)

// ограничить кол-во выполняемых do до 3

type Semaphore struct {
	sem chan struct{}
}

func NewSemaphore(n int) *Semaphore {
	return &Semaphore{
		sem: make(chan struct{}, n),
	}
}

func (s Semaphore) Aquire() {
	s.sem <- struct{}{}
}

func (s Semaphore) Release() {
	<-s.sem
}

func do(d time.Duration) {
	time.Sleep(d)

	fmt.Println("Done")
}

func main() {
	var arr []time.Duration

	sem := NewSemaphore(3)

	for i := 1; i <= 10; i++ {
		arr = append(arr, time.Duration(i*100)*time.Millisecond)
	}

	wg := &sync.WaitGroup{}

	for i, d := range arr {
		wg.Add(1)
		go func() {
			defer func() {
				sem.Release()
				wg.Done()
			}()
			sem.Aquire()
			fmt.Printf("goroutine %v start\n", i)

			do(d)
		}()
	}

	wg.Wait()
}
