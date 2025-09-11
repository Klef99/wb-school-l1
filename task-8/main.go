package main

import (
	"fmt"
)

func bitSwap(n int64, i int64, bit uint) (int64, error) {
	if !(bit == 0 || bit == 1) {
		return 0, fmt.Errorf("bit out of range")
	}

	// 5 2 1
	// 0101 5
	// 0010
	// 0111 -> 7
	if bit == 1 {
		return n | (1<<i - 1), nil
	}

	// 5 1 0
	// 0101 5
	// 1110 ^(0001)
	// 0100 -> 4
	return n &^ (1<<i - 1), nil
}

func main() {
	var n, i int64
	var bit uint

	fmt.Println("Input number, position (start from 1), bit")
	_, err := fmt.Scan(&n, &i, &bit)
	if err != nil {
		fmt.Println(err)
	}

	res, err := bitSwap(n, i, bit)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
}
