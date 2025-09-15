package main

import (
	"fmt"
)

func main() {
	var a, b int

	fmt.Scan(&a, &b)

	// Метод простой, но может привести к переполнению типа при слишком больших значениях
	a += b
	b = a - b
	a = a - b

	fmt.Println("First swap")
	fmt.Println(a, b)

	// Свап через XOR обмен более предпочтителен из-за отсутствия проблем с размерами типа int
	a ^= b
	b ^= a
	a ^= b

	fmt.Println("Second swap")
	fmt.Println(a, b)
}
