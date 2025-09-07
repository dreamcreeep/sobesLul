package main

import (
	"sync"
	"time"
)

// реализовать свой кэш
// у нас стала нагрузка 50/50 на чтение-запись, что будем делать?(написать решение с партициями)
// как бы тут сделал инвалидацию кэша? (ttl)
// стоит ли кэшировать запросы которых нет в базе и в кэше (404 ответы)

type ICache interface {
	Set(k, v string)
	Get(k string) (v string, ok bool)
}

// подумать над тем как заполнять партиции
// и какое кол-во партиций должно быть, к примеру на:
// RPS read = 80000
// RPS write = 5000
// От меня конкретно интервьюер ждал, что я посчитаю +\- память и как буду
// инициализировать все партиции, также спрашивал по способам инвалидации

type item struct {
	value      string
	expiration int64 // unix timestamp в наносекундах
}

type Cache struct {
	mu   sync.RWMutex
	data map[string]item
	ttl  time.Duration
}

func NewCache(ttl time.Duration, cleanupInterval time.Duration) *Cache {
	c := &Cache{
		data: make(map[string]item),
		ttl:  ttl,
	}

	go func() {
		ticker := time.NewTicker(cleanupInterval)
		defer ticker.Stop()

		for range ticker.C {
			c.cleanup()
		}
	}()

	return c
}

func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = item{
		value:      value,
		expiration: time.Now().Add(c.ttl).UnixNano(),
	}
}

func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	it, ok := c.data[key]
	if !ok || time.Now().UnixNano() > it.expiration {
		return "", false
	}
	return it.value, true
}

func (c *Cache) cleanup() {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now().UnixNano()
	for k, it := range c.data {
		if now > it.expiration {
			delete(c.data, k)
		}
	}
}

// type Partition struct {
//     data map[string]string
//     m sync.RWMutex
// }

// type Cache struct {
//     partitions []*Partition
// }

// func (c *Cache) Set(k, v string) {
//     partitionIndex := c.getPartitionIndex(k)
//     partition := c.partitions[partitionIndex]
//     partition.m.Lock()
//     defer partition.m.Unlock()
//     partition.data[k] = v
// }

// func (c *Cache) Get(k string) (string, bool) {
//     partitionIndex := c.getPartitionIndex(k)
//     partition := c.partitions[partitionIndex]
//     partition.m.RLock()
//     defer partition.m.RUnlock()
//     v, ok := partition.data[k]
//     return v, ok
// }

// func (c *Cache) getPartitionIndex(k string) int {
//     return hash % cap(c.partitions)
// }
