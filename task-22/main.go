package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := int64(5 << 40)
	b := int64(1 << 24)

	fmt.Println("a + b: ", a+b)
	fmt.Println("a * b: ", a*b) // = 0 - так как произошло переполнение инта
	fmt.Println("a / b: ", a/b)
	fmt.Println("a - b: ", a-b)

	ab := big.NewInt(a)
	bb := big.NewInt(b)
	z := bb.Mul(ab, bb)
	// Теперь результат будет нормальный
	fmt.Println("(big) a * b: ", z)
}
