package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
)

type Response struct {
	Err        error
	StatusCode int
	Url        string
}

// какие бы тесты тут написал?
// есть ли какая нибудь проблема со стандартным http.Get()?

func main() {
	var urls = []string{
		"http://ozon.ru",
		"http://sobakapodushka.ru",
		"https://vsepo20",
		"https://prikolnaprikole",
		"https://httpbin.org",
	}

	ctx := context.Background()

	respCh := make(chan Response)

	wg := &sync.WaitGroup{}

	for _, url := range urls {

		wg.Add(1)
		go func() {
			defer wg.Done()

			select {
			case <-ctx.Done():
				return
			default:
			}

			resp, err := http.Get(url)

			if err != nil {
				respCh <- Response{Err: err, Url: url}
			} else {
				respCh <- Response{StatusCode: resp.StatusCode, Url: url}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(respCh)
	}()

	for resp := range respCh {
		if resp.Err != nil || resp.StatusCode != 200 {
			fmt.Printf("адрес %s - not ok\n", resp.Url)
		} else {
			fmt.Printf("адрес %s - ok\n", resp.Url)
		}
	}
}
