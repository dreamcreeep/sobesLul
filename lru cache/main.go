package main

// реализовать LRU кэш

type cache interface {
	Set(k, v string)
	Get(k string) (v string, ok bool)
}

// type Node struct {
//     data string
//     next *Node
//     prev *Node
// }

// type Cache struct {
//     mu sync.RWMutex
//     data map[string]Node
//     head *Node
//     tail *Node
// }

// func New(limit int){
//     data := make(map[string]Node, limit)

//     return &Cache{
//         data: data,
//         limit: limit,
//     }
// }

// func (c *Cache) Set(k, v string){
//     c.mu.Lock
//     defer c.mu.Unlock

//     if len(c.data) == c.limit{
//         removeLast()
//     }

//     c.data[k] = v
//     v.next = c.head
//     c.head = v

//     if v.next == nil {
//         c.tail = c.head
//     }
// }

// func (c *Cache) Get(k string) (v Node, ok bool) {
//     c.mu.RWLock
//     defer c.mu.RWUnlock

//     // получаем ноду
//     v, ok := c.data[k]

//     if v == c.tail {
//         c.tail = v.prev
//         v.prev.next = nil
//     }

//     // сохр связ
//     v.prev.next = v.next
//     v.next.prev = v.prev

//     // для новой головы уст next
//     v.next = c.head
//     // уст новую голову
//     c.head = v

//     return v, ok
// }
