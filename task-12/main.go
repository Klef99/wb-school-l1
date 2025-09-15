package main

import (
	"fmt"
)

func set[T comparable](elems ...T) []T {
	res := make([]T, len(elems))
	data := make(map[T]struct{}, len(elems))

	for _, v := range elems {
		data[v] = struct{}{}
	}

	for k := range data {
		res = append(res, k)
	}

	return res
}

func main() {
	data := []string{"cat", "cat", "dog", "cat", "tree", "bush"}

	fmt.Println(set(data...))
}
