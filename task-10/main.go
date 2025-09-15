package main

import (
	"fmt"
)

func main() {
	res := make(map[int][]float64, 10)

	data := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	for _, v := range data {
		// получим значения десятков для определения группы
		d := 10 * int(v/10)
		// Несмотря на nil слайсы в мапе, операция append безопасна для использования.
		res[d] = append(res[d], v)
	}

	fmt.Println(res)
}
