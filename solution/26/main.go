/*
Разработать программу, которая проверяет, что все символы в строке уникальные
(true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.
Например:
abcd — true abCdefAaf — false aabcd — false
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(unique("abcd"))
	fmt.Println(unique("abCdefAaf"))
	fmt.Println(unique("aabcd"))
}

func unique(str string) bool {
	str = strings.ToLower(str)
	u := make(map[rune]struct{})
	for _, r := range str {
		if _, ok := u[r]; ok {
			return false
		}
		u[r] = struct{}{}
	}
	return true
}
