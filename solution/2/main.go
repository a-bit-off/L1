/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел
взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	fmt.Println("Squaring_Solution_1:", Squaring_Solution_1(numbers))
	fmt.Println("Squaring_Solution_2:", Squaring_Solution_2(numbers))
	fmt.Println("Squaring_Solution_3:", Squaring_Solution_3(numbers))
	fmt.Println("Squaring_Solution_4:", Squaring_Solution_3(numbers))
}

// Буферизованный канал
func Squaring_Solution_1(numbers []int) []int {
	length := len(numbers)
	squareNums := make([]int, length)
	ch := make(chan int, length)
	defer close(ch)

	// запускаем горутины в цикле и запысываем в канал данне
	for _, num := range numbers {
		go func(num int) {
			ch <- num * num
		}(num)
	}

	// считываем данные как только буфер заполнился
	for i := 0; i < length; i++ {
		square, ok := <-ch
		if !ok {
			return nil // обработка ошибки при чтении из канала
		}
		squareNums[i] = square
	}

	return squareNums
}

// Небуферизованный канал
func Squaring_Solution_2(numbers []int) []int {
	length := len(numbers)
	squareNums := make([]int, length)
	ch := make(chan int)
	defer close(ch)

	// запускаем горутины в цикле и запысываем в канал данные
	for _, num := range numbers {
		go func(num int) {
			ch <- num * num
		}(num)
	}

	// считываем данные как только буфер заполнился
	for i := 0; i < length; i++ {
		square, ok := <-ch
		if !ok {
			return nil // обработка ошибки при чтении из канала
		}
		squareNums[i] = square
	}

	return squareNums
}

// WaitGroup Mutex
func Squaring_Solution_3(numbers []int) []int {
	squareNums := make([]int, len(numbers))
	m := sync.Mutex{}
	wg := sync.WaitGroup{}

	for i, num := range numbers {
		wg.Add(1) // добавляем в счетчик +1(горутина)
		go func(num, i int, wg *sync.WaitGroup) {
			defer wg.Done() // по завершении горутины вычитаем из счетчика wg
			m.Lock()        // блокируем область памяти для предоставления ее горутине (уникальный доступ к области памяти)
			squareNums[i] = num * num
			m.Unlock()
		}(num, i, &wg)
	}

	wg.Wait()

	return squareNums
}

// WaitGroup одна горутина
func Squaring_Solution_4(numbers []int) []int {
	squareNums := make([]int, len(numbers))
	wg := sync.WaitGroup{}

	wg.Add(1) // добавляем в счетчик +1(горутина)
	go func(wg *sync.WaitGroup) {
		defer wg.Done() // по завершении горутины вычитаем из счетчика wg
		for i, num := range numbers {
			squareNums[i] = num * num
		}
	}(&wg)

	wg.Wait()

	return squareNums
}
