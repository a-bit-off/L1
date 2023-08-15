/*
Реализовать паттерн «адаптер» на любом примере.
*/
package main

import "fmt"

// Интерфейс, который описывает метод Request
type Target interface {
	Request() string
}

// Структура, которую мы хотим адаптировать
type Adaptee struct {
}

// Метод SpecificRequest, который мы хотим использовать через адаптер
func (a *Adaptee) SpecificRequest() string {
	return "Request"
}

// Адаптер, реализующий интерфейс Target
type Adapter struct {
	// первый вариант наследование
	*Adaptee
	// второй вариант объект
	//ad *Adaptee
}

// Реализация метода Request через вызов метода SpecificRequest адаптируемого объекта
func (a *Adapter) Request() string {
	return a.SpecificRequest()
}

func main() {
	adaptree := &Adaptee{}
	adapter := &Adapter{adaptree}

	fmt.Println(adapter.SpecificRequest())
	fmt.Println(adapter.Request())
}
