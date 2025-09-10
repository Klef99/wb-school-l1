package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Получаем количество секунд для работы из параметров
	n := flag.Int("n", 2, "Number of seconds")
	flag.Parse()

	data := make(chan int)

	// Горутина-читатель
	go func() {
		for d := range data {
			fmt.Println(d)
		}
	}()

	t := time.After(time.Duration(*n) * time.Second)

	// Пишем значения пока не получаем timeout
	for {
		select {
		case data <- rand.Intn(100):
			time.Sleep(100 * time.Millisecond)
		case <-t:
			fmt.Println("timeout")
			return
		}
	}
}
