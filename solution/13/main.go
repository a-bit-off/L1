/*
Поменять местами два числа без создания временной переменной.
*/

package main

import "fmt"

func main() {
	a := 1
	b := 2
	fmt.Println("a b")
	fmt.Println(a, b)

	fmt.Println(Solution1(a, b))
	fmt.Println(Solution2(a, b))
}

func Solution1(a, b any) (any, any) {
	a, b = b, a
	return a, b
}

func Solution2(a, b int) (int, int) {
	a = a + b
	b = a - b
	a = a - b
	return a, b
}
