package main

import (
	"math/rand"
	"sync"
)

func main() {
	// Из-за внутреннего устройства sync.Map при доминирующих операциях записи и удаления,
	// эта мапа будет медленнее и занимать больше памяти, чем обычная с sync.RWMutex/
	data := sync.Map{}
	wg := &sync.WaitGroup{}

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			data.Store(i, rand.Intn(100))
		}()
	}

	wg.Wait()
}

// Result
// time go run -race task-7/sync_map_version/main.go
// go run -race task-7/sync_map_version/main.go  0.12s user 0.19s system 16% cpu 1.831 total
