// Обьединить все каналы в 1
// Если 1 закрывается, все остальные закрыть
func merge(channels ...chan int) <-chan int {
	resultChan := make(chan int, len(channels))
	var wg sync.WaitGroup

	var once sync.Once

	ctx, cancel := context.WithCancel(context.Background())

	for _, channel := range channels {
		wg.Add(1)
		go func(channel chan int) {
			defer wg.Done()

			for {
				select {
				case val, ok := <-channel:
					if !ok {
						once.Do(cancel)
						return
					}

					select {
					case resultChan <- val:
					case <-ctx.Done():
						close(channel)
						return
					}

				case <-ctx.Done():
					close(channel)
					return
				}
			}
		}(channel)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	return resultChan
}
