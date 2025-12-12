package main

import (
	"fmt"
	"sync"
)

type WaitGroup struct {
	mu sync.RWMutex
}

func (w *WaitGroup) Go(f func()) {
	w.mu.RLock()

	go func() {
		f()
	}()
}

func (w *WaitGroup) Wait() {
	w.mu.Lock()
	// emty critical section
	w.mu.Unlock()
}

func main() {
	var wg WaitGroup

	for i := 0; i < 3; i++ {
		wg.Go(func() {
			fmt.Println("ok")
		})
	}
}
