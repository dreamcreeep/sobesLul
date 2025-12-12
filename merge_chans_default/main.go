package main

import (
	"context"
	"sync"
)

func merge(ctx context.Context, chans ...<-chan int) <-chan int {
	out := make(chan int)

	wg := &sync.WaitGroup{}

	for _, ch := range chans {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case val, ok := <-ch:
					if !ok {
						return
					}

					select {
					case <-ctx.Done():
						return
					case out <- val:
					}

				}
			}
		}()

	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
