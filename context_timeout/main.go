package main

import (
	"context"
	"fmt"
	"time"
)

const processingMaxTimeout = 3 * time.Second

func main() {
	var result int

	ctx, cancel := context.WithTimeout(context.Background(), processingMaxTimeout)
	defer cancel()

	// === edit here ===

	// =================

	fmt.Println(result)
}

func process() int {
	// long processing job
	time.Sleep(time.Second * 2)

	return 1
}
