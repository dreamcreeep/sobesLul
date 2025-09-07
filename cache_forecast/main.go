package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

type Cache struct {
	info     int
	interval time.Duration
	mtx      sync.RWMutex
}

func NewCache(interval time.Duration) *Cache {
	info := aiWeatherForecast()
	return &Cache{
		info:     info,
		interval: interval,
	}
}

func (c *Cache) Get() int {
	c.mtx.RLock()
	defer c.mtx.RUnlock()

	return c.info
}

func (c *Cache) Update(ctx context.Context) {
	ticker := time.NewTicker(c.interval * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			updInfo := aiWeatherForecast()
			c.mtx.Lock()
			c.info = updInfo
			c.mtx.Unlock()
		}
	}
}

// aiWeatherForecast через нейронную сеть вычисляет прогноз погоды за ~1 секунду
func aiWeatherForecast() int {
	time.Sleep(1 * time.Second)
	return rand.Intn(70) - 30
}

func main() {
	ctx := context.Background()

	cache := NewCache(3)

	go cache.Update(ctx)

	http.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "{\"temperature\":%d}\n", cache.Get())
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
