/*
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
Приведите корректный пример реализации.

var justString string

	func someFunc() {
		v := createHugeString(1 << 10)
		justString = v[:100]
	}

	func main() {
		someFunc()
	}
*/
package main

import (
	"fmt"
	"log"
)

// 1. Использование глобальной переменной плохая практика
var justString string

func someFunc() {
	// 2. Размер стоки в слайсе и функции createHugeString
	// задается колличеством байт, плохая практика при работе с string

	// 3. Из 1024 используется только первые 100,
	// можно очистить лишнее самотоятельно runtime.KeepAlive(v)
	// либо создавать соизмерно необходимому

	v := createHugeString(1 << 10)
	justString = v[:100] // 4. Возможно словить панику out of range
}

func createHugeString(size int) (res string) {
	if size < 0 {
		log.Fatalln("Invalid size!")
	}

	for i := 0; i < size; i++ {
		res += string(i + 18000)
	}
	return res

}

func main() {
	someFunc()
	fmt.Println(justString)

	res, err := bestSomeFunc()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)

}

func bestSomeFunc() (string, error) {
	res := string(make([]byte, 100))

	v := createHugeString(1 << 10)
	if len(v) >= 100 {
		res = v[:100]
	} else {
		return res, fmt.Errorf("Out of range!")
	}

	return res, nil
}
