package main

import (
	"math/rand"
	"strings"
)

var justString string

func createHugeString(size int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	n := len(letters)
	bld := strings.Builder{}

	for i := 0; i < size; i++ {
		bld.WriteRune(letters[rand.Intn(n)])
	}

	return bld.String()
}

// 1. Переменная justString является строкой, основанной на срезе по большой строке.
// 2. Так как переменная глобальная, её время жизни больше, чем у someFunc
// 3. Срез по большой строке не позволят GC высвободить лишнюю память.
// 4. Большое количество вызов этой функции будет приводить к излишнему расходу памяти.
func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

// Решение состоит в копировании строки, а не создании среза на уже имеющейся.
// Также можно вернуть строку из функции, а не использовать глобальную переменную.
func someFuncCorrect() {
	v := createHugeString(1 << 10)

	elem := make([]byte, 100)
	copy(elem, v)

	justString = string(elem)
}

func main() {
	someFunc()
}
