package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	data := [5]int{2, 4, 6, 8, 10}

	wg := &sync.WaitGroup{}
	wg.Add(len(data))

	for _, d := range data {
		go func() {
			defer wg.Done()
			fmt.Println(math.Pow(float64(d), 2))
		}()
	}

	wg.Wait()
}
