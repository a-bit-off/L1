/*
Дана последовательность температурных колебаний: -25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/
package main

import (
	"fmt"
)

func main() {
	temperatures := []float64{
		-31.1,
		-29.9, -20.0,
		-19.9, -15.5, -10.0,
		-9.9, -5.5, 0, 3, 7, 9,
		10, 15, 19,
		20, 25,
		30,
	}
	fmt.Println(Solution1(temperatures))
}

func Solution1(temperatures []float64) map[int][]float64 {
	// создаем мапу с ключом int (group)
	// и значением массив не целых чисел []float64 (groups temp)
	groupedTemp := make(map[int][]float64)

	for _, temp := range temperatures {
		group := int(temp/10) * 10                            // получаем группу (-12.5 / 10) * 10 = -10
		groupedTemp[group] = append(groupedTemp[group], temp) // добавляем в группу соответствующую ей температуру
	}
	return groupedTemp
}
