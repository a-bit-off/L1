/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/
package main

import "fmt"

func main() {
	nums := []int{3, 2, 1, 5, 8, 6, 5}
	res := quickSort(nums)
	fmt.Println(res)
}

func quickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	pivot := nums[0]
	left := make([]int, 0)
	right := make([]int, 0)

	for _, num := range nums[1:] {
		if num <= pivot {
			left = append(left, num) // если число меньше чем pivot добавляем в левую группу
		} else {
			right = append(right, num) // иначе в правую
		}
	}

	// рекурсивно возвращаем все значения
	// соединяем левую, опорную и правую части
	return append(append(quickSort(left), pivot), quickSort(right)...)
}
