/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/
package main

import "fmt"

type Human struct {
	name string
	age int
}

type Action struct {
	Human // встраиваем структуру
}

func (h Human) Greet() {
	fmt.Printf("Hi, my name is %s, i am %d years old\n", h.name, h.age)
}

func (a *Action)HaveBirthday() {
	a.age++
}

func main() {
	human := Human{name: "Bob", age: 7}
	action := Action{human}
	action.Greet()
	action.HaveBirthday()
	action.Greet()
}