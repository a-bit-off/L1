/*
Дана последовательность температурных колебаний: -25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 5, 6, -1, -11, 1, -29.9, 32.5}
	groupedTemp := make(map[int][]float64)

	for _, temp := range temperatures {
		group := int(math.Ceil(temp/10) * 10)
		groupedTemp[group] = append(groupedTemp[group], temp)
	}

	fmt.Println(groupedTemp)
}
