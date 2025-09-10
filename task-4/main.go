package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context, id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Goroutine %d: done\n", id)
			return
		case job := <-jobs:
			fmt.Printf("Goroutine %d: %d\n", id, job)
		}
	}
}

func main() {
	// Получение параметра при запуске
	n := flag.Int("n", 3, "number of workers")
	flag.Parse()

	jobs := make(chan int)
	quit := make(chan os.Signal, 1)
	wg := &sync.WaitGroup{}

	// Context выбран вместо канала для завершения работы воркеров,
	//так как контекст в дальнейшем может использоваться в других целях (например трассировка),
	//а не только для graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(*n)
	for i := 0; i < *n; i++ {
		go worker(ctx, i, jobs, wg)
	}

	// Подписываемся на сигнал в канал quit для graceful shutdown
	signal.Notify(quit, syscall.SIGINT)

	for {
		select {
		case <-quit:
			cancel()
			wg.Wait()
			fmt.Println("All workers done")
			return
		case jobs <- rand.Intn(100):
			time.Sleep(100 * time.Millisecond) // задержка для наглядности записи и чтения
		}
	}
}
