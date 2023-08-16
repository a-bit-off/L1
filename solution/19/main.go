/*
Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode.
*/

package main

import (
	"fmt"
)

func main() {
	str := "гла䙐врыба䨸"
	fmt.Println(str)

	fmt.Println(solution1(str))
	fmt.Println(solution2(str))
}

func solution1(str string) string {
	n := len(str)
	reversed := make([]rune, n)
	for i, r := range str {
		reversed[n-1-i] = r // 10-1-0 = r; 10-1-1 = r таким образом обеспечиватеся переворот
	}
	return string(reversed)
}

// наименее производительное решение
func solution2(str string) (reversed string) {
	for _, r := range str {
		reversed = string(r) + reversed // a; ba; cba; к началу добавлям предыдущую строку
	}
	return reversed
}
