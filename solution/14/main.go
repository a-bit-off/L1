/*
Разработать программу, которая в рантайме способна определить тип переменной:
int, string, bool, channel из переменной типа interface{}.
*/
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var variable interface{}

	variable = 21
	GetType(variable)

	variable = "21"
	GetType(variable)

	variable = true
	GetType(variable)

	variable = make(chan int)
	GetType(variable)

	variable = 45.5
	GetType(variable)
}

func GetType(variable interface{}) {
	switch reflect.TypeOf(variable).Kind() {
	case reflect.Int:
		fmt.Println("Тип: int")
	case reflect.String:
		fmt.Println("Тип: string")
	case reflect.Bool:
		fmt.Println("Тип: bool")
	case reflect.Chan:
		fmt.Println("Тип: chan")
	default:
		fmt.Println("Тип: неизвестный")
	}
}
