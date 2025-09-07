package main

import (
	"fmt"
	"runtime"
	"time"
)

// выведется ли done?

func main() {
	runtime.GOMAXPROCS(1)

	var n int

	go func() {
		for {
			n++
		}
	}()

	time.Sleep(500 * time.Millisecond)

	fmt.Println("Done")
}
