/*
Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
*/
package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func main() {
	p1 := NewPoint(8, 10)
	p2 := NewPoint(3, 5)

	fmt.Println(distance1(p1, p2))
	fmt.Println(distance2(p1, p2))
}

func distance1(p1, p2 *Point) float64 {
	// формула расчета расстояния на коор плоскости
	a := p1.x - p2.x
	b := p1.y - p2.y
	return math.Sqrt(a*a + b*b)
}

// менее производительное решение
func distance2(p1, p2 *Point) float64 {
	a := math.Pow(p1.x-p2.x, 2)
	b := math.Pow(p1.y-p2.y, 2)
	d := math.Sqrt(a + b)
	return d
}
