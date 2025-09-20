package main

import (
	"fmt"
)

func delElem[T any](a []T, i int) ([]T, bool) {
	if i >= len(a) {
		return []T{}, false
	}

	copy(a[i:], a[i+1:])
	clear(a[len(a)-1:]) // очистить элемент в конце (zero/nil в зависимости от типа)

	return a[:len(a)-1], true
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("Original: ", a, "Len: ", len(a), "Cap: ", cap(a))
	a, ok := delElem(a, 6)
	if ok {
		fmt.Println("Delete: ", a, "Len: ", len(a), "Cap: ", cap(a))
	}
}
