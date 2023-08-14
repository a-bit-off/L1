/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
*/

package main

import (
	"fmt"
)
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 56}

	go func(ch1 chan int, nums []int) {
		defer close(ch1)

		for _, num := range nums {
			ch1 <- num
		}
	}(ch1, nums)

	go func(ch1, ch2 chan int, nums []int) {
		defer close(ch2)
		for num := range ch1 {
			ch2 <- num * 2
		}

	}(ch1, ch2, nums)


	for num := range ch2 {
		fmt.Println(num)
	}
}

