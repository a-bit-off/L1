/*
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….)
с использованием конкурентных вычислений.
*/
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	fmt.Println("SquaringSum_Solution_1:", SquaringSum_Solution_1(numbers))
	fmt.Println("SquaringSum_Solution_2:", SquaringSum_Solution_2(numbers))
	fmt.Println("SquaringSum_Solution_3:", SquaringSum_Solution_3(numbers))
	fmt.Println("SquaringSum_Solution_4:", SquaringSum_Solution_4(numbers))
}

// atomic
func SquaringSum_Solution_1(numbers []int) int {
	done := make(chan struct{})
	defer close(done)
	var squaringSum atomic.Int32
	length := len(numbers)

	for _, num := range numbers {
		go func(num int) {
			// атомарное выполнение программы
			squaringSum.Add(int32(num * num))
			done <- struct{}{}
		}(num)
	}

	for i := 0; i < length; i++ {
		<-done
	}

	return int(squaringSum.Load())
}

// WaitGroup Mutex
func SquaringSum_Solution_2(numbers []int) int {
	var squaringSum int
	m := sync.Mutex{}
	wg := sync.WaitGroup{}

	for _, num := range numbers {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			m.Lock()
			squaringSum += num * num
			m.Unlock()
		}(num)
	}

	wg.Wait()

	return squaringSum
}

// Небуферизованный канал
func SquaringSum_Solution_3(numbers []int) int {
	length := len(numbers)
	var squaringSum int
	ch := make(chan int)
	defer close(ch)

	for _, num := range numbers {
		go func(num int) {
			ch <- num * num
		}(num)
	}

	// Ждем завершения всех горутин и закрываем канал после этого
	for i := 0; i < length; i++ {
		square, ok := <-ch
		if !ok {
			return 0 // обработка ошибки при чтении из канала
		}
		squaringSum += square
	}

	return squaringSum
}

// Буферизованный канал
func SquaringSum_Solution_4(numbers []int) int {
	var squaringSum int
	length := len(numbers)
	ch := make(chan int, length)
	defer close(ch)

	for _, num := range numbers {
		go func(num int) {
			ch <- num * num
		}(num)
	}

	// Ждем завершения всех горутин и закрываем канал после этого
	for i := 0; i < length; i++ {
		square, ok := <-ch
		if !ok {
			return 0 // обработка ошибки при чтении из канала
		}
		squaringSum += square
	}

	return squaringSum
}
