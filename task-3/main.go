package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func worker(id int, jobs <-chan int) {
	for job := range jobs {
		fmt.Printf("Goroutine %d: %d\n", id, job)
	}
}

func main() {
	// Получение параметра при запуске
	n := flag.Int("n", 3, "number of workers")
	flag.Parse()

	jobs := make(chan int)

	for i := 0; i < *n; i++ {
		go worker(i, jobs)
	}

	for {
		jobs <- rand.Intn(100)
		time.Sleep(100 * time.Millisecond) // задержка для наглядности записи и чтения
	}
}
