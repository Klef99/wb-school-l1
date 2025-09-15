package main

import (
	"fmt"
)

func main() {
	in := make(chan int)
	out := make(chan int)

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	go func() {
		for _, v := range data {
			in <- v
		}
		close(in)
	}()

	go func() {
		for d := range in {
			out <- d * d
		}
		close(out)
	}()

	for v := range out {
		fmt.Println(v)
	}
}
