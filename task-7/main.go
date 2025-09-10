package main

import (
	"math/rand"
	"sync"
)

type SafeMap struct {
	mu *sync.RWMutex
	m  map[int]int
}

// NewSafeMap Конструктор для конкурентной мапы. n - для предварительной аллокации достаточного места
func NewSafeMap(n int) *SafeMap {
	return &SafeMap{
		mu: &sync.RWMutex{},
		m:  make(map[int]int, n),
	}
}

func (s *SafeMap) Set(key int, value int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[key] = value
}

func (s *SafeMap) Get(key int) int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.m[key]
}

func main() {
	mp := NewSafeMap(100)
	wg := &sync.WaitGroup{}

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			mp.Set(i, rand.Intn(100))
		}()
	}

	wg.Wait()
}

// Result
// time go run -race task-7/main.go
// go run -race task-7/main.go  0.10s user 0.14s system 13% cpu 1.716 total
