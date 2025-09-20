package main

import (
	"bufio"
	"fmt"
	"os"
)

// Разворот одного слова на месте. Границы - [l,r)
func reverseWord(s []rune, l, r int) {
	i := l
	j := r - 1

	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

func reverseWords(s []rune) {
	// Разворот строки полностью
	l := 0
	r := len(s) - 1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}

	l = 0
	r = 0
	// Разворачиваем слова
	for r < len(s) {
		for r < len(s) && s[r] != ' ' {
			r++
		}

		reverseWord(s, l, r)
		r++
		l = r
	}
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	data := []rune(sc.Text())

	fmt.Println("Origin: ", string(data))
	reverseWords(data)
	fmt.Println("Reversed: ", string(data))
}
