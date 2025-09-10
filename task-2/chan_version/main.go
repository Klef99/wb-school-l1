package main

import (
	"fmt"
)

func worker(data <-chan int, results chan<- int) {
	n := <-data
	results <- n * n
}

func main() {
	source := [5]int{2, 4, 6, 8, 10}

	data := make(chan int)
	results := make(chan int)

	for _ = range len(source) {
		go worker(data, results)
	}

	go func() {
		for _, n := range source {
			data <- n
		}

		close(data)
	}()

	for _ = range len(source) {
		fmt.Println(<-results)
	}
}
