package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func intersections[T comparable](arr1 []T, arr2 []T) []T {
	ml := min(len(arr1), len(arr2))
	res := make([]T, 0, ml)
	tmp := make(map[T]struct{}, ml)

	for _, v := range arr1 { // добавим элементы из первого слайса в мапу, для обнаружения повторов.
		tmp[v] = struct{}{}
	}

	for _, v := range arr2 {
		if _, ok := tmp[v]; ok {
			res = append(res, v)
			delete(tmp, v)
		}
	}

	return res
}

func main() {
	var s1, s2 string

	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < 2 && scanner.Scan(); i++ {
		switch {
		case i == 0:
			s1 = scanner.Text()
		case i == 1:
			s2 = scanner.Text()
		}
	}

	set1 := strings.Split(s1, ", ")
	set2 := strings.Split(s2, ", ")

	fmt.Println(intersections(set1, set2))
}
