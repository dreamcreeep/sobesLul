package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// 1. конкурентно по батчам запросить данные и записать в файл
// 2. сделать так, чтобы одновременно выполнялось не более chunkSize запросов

const (
	url        = `http://jsonplaceholder.typicode.com/tools/%d`
	chunkSize  = 100
	dataCount  = 2 << 10
	numWorkers = 4
	numPages   = 100
)

func proceed() {

	for j := 0; j < numWorkers; j++ {
		go func() {
			k := j
			for k < numPages {
				url := fmt.Sprintf(url, k)
				resp, err := http.Get(url)
				k += numWorkers
			}
		}()
	}

	batch := make([]byte, 0, dataCount)

	go func() {
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()

		file, err := os.OpenFile("aafaf.html", 2, 0666)
		if err != nil {
			fmt.Printf("%v", err)
		}

		defer file.Close()

		if _, err := io.Copy(file, resp.Body); err != nil {
			fmt.Printf("%v", err)
		}
	}()

}
