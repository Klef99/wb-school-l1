package main

import (
	"fmt"
	"strings"
)

func uniqueString(s string) bool {
	data := []rune(strings.ToLower(s))
	cnt := make(map[rune]int, len(data)) // Воспользуемся мапой для проверки на уникальность

	for _, v := range data {
		if _, ok := cnt[v]; ok {
			return false // Если ключ уже в мапе - руна повторилась
		}
		cnt[v]++
	}

	return true
}

func main() {
	fmt.Println("abcd: ", uniqueString("abcd"))
	fmt.Println("abCdefAaf: ", uniqueString("abCdefAaf"))
	fmt.Println("aabcd: ", uniqueString("aabcd"))
}
