// при классческой реализации мы создаем на каждый канал по горутине, но хотели бы контролировать их количество и создавать определенное
// реализована идея с воркерами, есть общий канал с тасками, куди им пишут значения,  а они читают и обрабатывают, таски это поданые каналы

package main

import "sync"

const numWorker = 10

func merge[T any](channels ...<-chan T) <-chan T {
	ch := make(chan T)
	wg := sync.WaitGroup{}

	tasks := make(chan (<-chan T), len(channels))
	// для того, чтобы потом закрыть tasks и выйти из горутин
	openChannelsWg := sync.WaitGroup{}

	for _, channel := range channels {
		tasks <- channel
	}
	openChannelsWg.Add(len(channels))

	for range numWorker {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for channel := range tasks {
				select {
				case v, ok := <-channel:
					if !ok {
						// если канал закрыт вычитаем его из открытых
						// и идём дальше
						openChannelsWg.Done()
						continue
					}
					ch <- v
				default: // нет данных
				}

				// возвращаем канал обратно, чтобы потом снова обрабатывать
				tasks <- channel
			}
		}()
	}

	go func() {
		// сначала ждём, чтобы все каналы закрылись
		openChannelsWg.Wait()
		// после закрываем таски, что вызовет возвращение горутин
		close(tasks)
		wg.Wait()
		close(ch)
	}()

	return ch
}
