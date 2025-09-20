package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomSlice(n int) []int {
	res := make([]int, 0, n)
	for i := 0; i < n; i++ {
		res = append(res, rand.Intn(10*n))
	}

	return res
}

// Рекурсивный алгоритм с дополнительной аллокацией слайсов.
func quickSortWithAlloc(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	reference := len(arr) / 2
	left := make([]int, 0, len(arr))
	eq := make([]int, 0, len(arr))
	right := make([]int, 0, len(arr))

	for i, elem := range arr {
		switch {
		case elem < arr[reference]:
			left = append(left, arr[i])
		case elem == arr[reference]:
			eq = append(eq, arr[i])
		case elem > arr[reference]:
			right = append(right, arr[i])

		}
	}

	left = append(quickSortWithAlloc(left), eq...)

	return append(left, quickSortWithAlloc(right)...)
}

func partition(arr []int, l, r int) int {
	v := arr[(l+r)/2]
	i := l
	j := r

	for i <= j {
		for arr[i] < v {
			i++
		}
		for arr[j] > v {
			j--
		}
		if i >= j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}

	return j
}

// Рекурсивный алгоритм, работающий на исходном слайсе.
func quickSortInplace(arr []int, l, r int) {
	if l < r {
		p := partition(arr, l, r)

		quickSortInplace(arr, l, p)
		quickSortInplace(arr, p+1, r)
	}
}

func main() {
	arr := randomSlice(10)
	fmt.Printf("Original array: %v\n", arr)

	now := time.Now()
	arrN := quickSortWithAlloc(arr)
	fmt.Printf("With allocations (elapsed: %v): %v\n", time.Since(now), arrN)

	now = time.Now()
	quickSortInplace(arr, 0, len(arr)-1)
	fmt.Printf("Inplace (elapsed: %v): %v\n", time.Since(now), arrN)
}
