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
}

// WaitGroup Mutex
func Squaring_Solution_1(numbers []int) []int {
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

// Небуферизованный канал
func Squaring_Solution_2(numbers []int) []int {
	length := len(numbers)
	squareNums := make([]int, length)
	ch := make(chan int)
	done := make(chan struct{})

	// запускаем горутину в котором в цикле будут записыватся данные в канал ch
	// done будет сигнализировать о том, что горутина завершилась
	go func() {
		defer close(ch)
		defer close(done)

		for _, num := range numbers {
			ch <- num * num
		}
	}()

	for i := 0; i < length; i++ {
		select {
		case num := <-ch: // пока цикл работает мы получаем данные
			squareNums[i] = num
		case <-done: // после завершения горутины мы получим сигнал
			return squareNums
		}
	}

	return squareNums
}

// Буферизованный канал
func Squaring_Solution_3(numbers []int) []int {
	length := len(numbers)
	squareNums := make([]int, length)
	ch := make(chan int, length)

	// создаем горутину в которой будем заполнять буф канал, до тех пор пока он не наполнится
	go func() {
		defer close(ch)

		for _, num := range numbers {
			ch <- num * num
		}
	}()

	// цикл начнет считывать с канал только при условию наполнения канала
	// разом считываем все данные
	for i := 0; i < length; i++ {
		squareNums[i] = <-ch
	}

	return squareNums
}
