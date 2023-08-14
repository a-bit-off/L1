/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "11 22 33 44 55 66"
	fmt.Println(solution1(str))
}

func solution1(str string) string {
	words := strings.Fields(str)
	size := len(words)
	mid := size / 2

	for i := 0; i < mid; i++ {
		j := size - 1 - i
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}
