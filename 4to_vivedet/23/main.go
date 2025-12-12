package main

import (
	"fmt"
	"sync"
	"time"
)

func testGoroutines() {
	dataMap := make(map[string]int)

	mtx := sync.Mutex{}

	for i := 0; i < 10000; i++ {
		go func(d map[string]int, num int) {
			mtx.Lock()
			d[fmt.Sprintf("%d", num)] = num
			mtx.Unlock()
		}(dataMap, i)
	}

	time.Sleep(5 * time.Second)
	fmt.Println(len(dataMap))
}

func main() {
	testGoroutines()
}
