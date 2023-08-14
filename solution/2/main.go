/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел
взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/
package main

import(
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	fmt.Println("Squaring_Solution_1:", Squaring_Solution_1(numbers))
	fmt.Println("Squaring_Solution_2:", Squaring_Solution_2(numbers))
	fmt.Println("Squaring_Solution_3:", Squaring_Solution_3(numbers))
}

// WaitGroup Mutex
func Squaring_Solution_1(numbers []int) []int {
	squareNums := make([]int, 0, len(numbers))
	m := sync.Mutex{}
	wg := sync.WaitGroup{}

	for _, num := range numbers {
		wg.Add(1)
		go func (num int) {
			defer wg.Done()
			m.Lock()
			squareNums = append(squareNums, num * num)
			m.Unlock()
		}(num)
	}

	wg.Wait()

	return squareNums
}

// Небуферизованный канал
func Squaring_Solution_2(numbers []int) []int {
	length := len(numbers)
	squareNums := make([]int, 0, length)
	ch := make(chan int)
	defer close(ch)

	for _, num := range numbers {
		go func(num int) {
			ch <- num * num
		}(num)
	}

	for i := 0; i < length; i++ {
		squareNums = append(squareNums, <-ch)
	}

	return squareNums
}

// Буферизованный канал
func Squaring_Solution_3(numbers []int) []int {
	length := len(numbers)
	squareNums := make([]int, 0, length)
	ch := make(chan int, length)
	defer close(ch)

	for _, num := range numbers {
		go func(num int) {
			ch <- num * num
		}(num)
	}

	for i := 0; i < length; i++ {
		squareNums = append(squareNums, <-ch)
	}

	return squareNums
}