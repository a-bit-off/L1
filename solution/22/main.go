/*
Разработать программу, которая перемножает, делит, складывает,
вычитает две числовых переменных a,b, значение которых > 2^20.
*/

package main

import (
	"fmt"
	"math/big"
)

func main() {
	num1 := big.NewInt(1 << 40)
	num2 := big.NewInt(1 << 30)

	sum := new(big.Int).Add(num1, num2)
	sub := new(big.Int).Sub(num1, num2)
	mul := new(big.Int).Mul(num1, num2)
	div := new(big.Int).Div(num1, num2)

	fmt.Println("num1:", num1)
	fmt.Println("num2:", num2)
	fmt.Println("sum:", sum)
	fmt.Println("sub:", sub)
	fmt.Println("mul:", mul)
	fmt.Println("div:", div)
}
