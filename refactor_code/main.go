package main

import "sync"

type Set[T any] struct {
	m  map[T]struct{}
	mu sync.Mutex
}

func New[T any]() Set[T] {
	return Set[T]{
		m: make(map[T]struct{}),
	}
}

func (s Set[T]) Has(k T) bool {
	var rw sync.RWMutex
	rw.RLock()
	defer rw.RUnlock()
	_, ok := s.m[k]
	return ok
}

func (s Set[T]) Set(k T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[k] = struct{}{}
}

func (s Set[T]) Intersection(o Set[T]) map[T]struct{} {
	var res map[T]struct{}
	for k, _ := range s {
		if o.Has(k) {
			res[k] = struct{}{}
		}
	}
	return res
}
