package main

import (
	"bufio"
	"fmt"
	"os"
)

// Разворачивает полностью слайс.
func reverse[T any](arr []T) []T {
	res := make([]T, 0, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		res = append(res, arr[i])
	}

	return res
}

func main() {
	var data string

	scn := bufio.NewScanner(os.Stdin)
	scn.Scan()
	data = scn.Text() // Сканируем полную строку.

	res := string(reverse([]rune(data)))

	fmt.Println("Origin: ", data)
	fmt.Println("Reversed: ", res)

}
