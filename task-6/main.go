package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func stopByCondition() {
	i := 0
	for {
		if i == 5 {
			fmt.Println("stopByCondition")
			return
		}
		fmt.Printf("condition goroutine work: %d\n", i)
		i++
		time.Sleep(100 * time.Millisecond)
	}
}

func stopByChan(done <-chan struct{}) {
	for {
		select {
		case <-done:
			fmt.Println("stopByChan")
			return
		default:
			fmt.Println("chan goroutine working")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func stopByClosedChan(job <-chan struct{}) {
	defer fmt.Println("stopByClosedChan")
	for range job {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("closed chan goroutine work")
	}
}

func stopByCtx(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stopByCtx")
			return
		default:
			fmt.Println("context goroutine working")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func stopByGoexit() {
	defer fmt.Println("stopByGoexit")

	time.Sleep(100 * time.Millisecond)
	fmt.Println("goexit goroutine working")

	runtime.Goexit()
	fmt.Println("never printed")
}

func stopByTimeout(timer *time.Timer) {
	for {
		select {
		case <-timer.C:
			fmt.Println("stopByTimeout")
			return
		default:
			fmt.Println("timeout goroutine working")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func main() {
	// Остановка по условию
	go stopByCondition()

	time.Sleep(600 * time.Millisecond)

	// Остановка по сигналу в канал
	done := make(chan struct{})

	go stopByChan(done)

	time.Sleep(500 * time.Millisecond)
	done <- struct{}{}

	// Остановка после закрытия канала
	go stopByClosedChan(done)
	for i := 0; i < 5; i++ {
		done <- struct{}{}
	}

	time.Sleep(600 * time.Millisecond)
	close(done)

	// Остановка через контекст
	// Условия могут быть разные (по timeout, deadline, через функцию остановки)
	// Удобен для остановки большого количества горунтин
	ctx, cancel := context.WithCancel(context.Background())

	go stopByCtx(ctx)

	time.Sleep(500 * time.Millisecond)
	cancel()

	// Остановка по runtime.Goexit()
	// Выполняет все defer и останавливает горутину
	// При наличии recover в defer они возвращают nil, так как это не является паникой
	go stopByGoexit()

	time.Sleep(500 * time.Millisecond)

	// Остановка по timeout. Можно выполнят как с помощью таймеров, так и контекстов с timeout
	timer := time.NewTimer(400 * time.Millisecond)

	go stopByTimeout(timer)

	time.Sleep(500 * time.Millisecond)
	timer.Stop()
}
