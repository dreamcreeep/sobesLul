package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	_, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	time.Sleep(1 * time.Second)

	cancel()
	fmt.Println("cancelled with cancel")

}
