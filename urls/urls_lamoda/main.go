package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Требуется обойти все url и получить слайс с кодами ответов в том же порядке
// Хотим делать параллельно но не больше k запросов одновременно

var urls = []string{
	"https://www.lamoda.ru/p/mp002xw0lvkd/clothes-tomollyfromjames-plate/",
	"https://www.lamoda.ru/p/mp002xw14uf2/clothes-tomollyfromjames-plate/",
	"https://www.lamoda.ru/p/rtladr746901/clothes-iceberg-plate/",
	"https://www.lamoda.ru/p/mp002xw18h9d/clothes-victoriaveisbrut-plate/",
	"https://www.lamoda.ru/p/mp002xw004x4/clothes-clanvi-plate/",
	"https://www.lamoda.ru/p/mp002xw0zfxy/clothes-glvr-plate/",
	"https://www.lamoda.ru/p/mp002xw0slmg/clothes-snezhnayakoroleva-plate-kozhanoe/",
	"https://www.lamoda.ru/p/mp002xw132c3/clothes-auranna-plate/",
	// ......
}

func crawl(urls []string, k int) []int {
	n := len(urls)
	res := make([]int, n)
	sem := make(chan struct{}, k)
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i, url := range urls {

		wg.Add(1)
		go func(i int, url string) {
			defer wg.Done()

			sem <- struct{}{}
			defer func() {
				<-sem
			}()

			client := http.Client{
				Timeout: 5 * time.Second,
			}

			resp, err := client.Get(url)

			if err != nil {

				res[i] = 0
				return
			}

			mu.Lock()
			res[i] = resp.StatusCode
			mu.Unlock()

			res.Body.Close()
		}(i, url)

	}
	wg.Wait()
	return res
}

func main() {
	result := crawl(urls, 5)
	fmt.Println("All done")
}
