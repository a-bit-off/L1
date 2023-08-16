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
	words := strings.Fields(str) // получаем слайс и слов
	size := len(words)
	mid := size / 2
	j := 0

	for i := 0; i < mid; i++ {
		j = size - 1 - i                        // 10-1-0
		words[i], words[j] = words[j], words[i] // заменяем значения в слайсе по индексу
	}
	return strings.Join(words, " ")
}
