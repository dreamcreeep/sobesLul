package main

import (
	"context"
	"fmt"
	"net/http"
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

	for _, url := range urls {
		go func(url string) {
			err := fetch(context.Background(), url)
			if err != nil {
				fmt.Printf("err %s\n", err)
				return
			}
		}(url)
	}
	fmt.Println("All requests launched!")
	time.Sleep(400 * time.Millisecond)
	fmt.Println("Done")
}

func fetch(ctx context.Context, url string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	_, err = http.DefaultClient.Do(req)
	return err
}
