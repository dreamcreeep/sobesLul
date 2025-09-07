package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// исправить код
// сделать отмену запросов после ошибки

func main() {
	urls := []string{
		"https://lamoda.ru",
		"https://yandex.ru",
		"http://mail.ru",
		"https://ya.ru",
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // обязательно, чтобы освободить ресурсы контекста

	wg := &sync.WaitGroup{}

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			err := fetch(ctx, url)
			if err != nil {
				fmt.Printf("err: %s (%s)\n", err, url)
				cancel() // отменим все остальные
				return
			}
			fmt.Printf("ok: %s\n", url)
		}(url)
	}

	fmt.Println("All requests launched!")
	wg.Wait()
	fmt.Println("Done")
}

func fetch(ctx context.Context, url string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	client := &http.Client{
		Timeout: 3 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	return nil
}
