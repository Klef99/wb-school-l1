package main

import (
	"fmt"
	"time"
)

// Реализация sleep через таймер
func sleepTimer(duration time.Duration) {
	<-time.After(duration)
}

// Проверка времени через цикл
func sleepFor(duration time.Duration) {
	deadline := time.Now().Add(duration)
	for !time.Now().After(deadline) {
	}
}

// Проверка через закрытие канала (схоже по сути с таймером)
func sleepChan(duration time.Duration) {
	done := make(chan struct{})
	time.AfterFunc(
		duration, func() {
			close(done)
		},
	)

	<-done
}

func main() {
	// Timer
	now := time.Now()
	sleepTimer(500 * time.Millisecond)
	fmt.Printf("elapsed (timer): %s\n", time.Since(now))

	// For
	now = time.Now()
	sleepFor(500 * time.Millisecond)
	fmt.Printf("elapsed (for): %s\n", time.Since(now))

	// Chan
	now = time.Now()
	sleepChan(500 * time.Millisecond)
	fmt.Printf("elapsed (chan): %s\n", time.Since(now))

}
