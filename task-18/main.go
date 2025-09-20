package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	cnt atomic.Int64
}

func NewCounter() *Counter {
	return &Counter{cnt: atomic.Int64{}}
}

func (c *Counter) Add(delta int64) {
	c.cnt.Add(delta)
}

func (c *Counter) Load() int64 {
	return c.cnt.Load()
}

func main() {
	const n = 10

	cnt := NewCounter()

	wg := &sync.WaitGroup{}
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			cnt.Add(1)
		}()
	}

	wg.Wait()

	fmt.Printf("Want: %d Get: %d\n", n, cnt.Load())
}
