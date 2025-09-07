package main

import "sync"

func merge(chans ...<-chan int) <-chan int {
	out := make(chan int)

	wg := &sync.WaitGroup{}

	for _, ch := range chans {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for val := range ch {
				out <- val
			}
		}()

	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
