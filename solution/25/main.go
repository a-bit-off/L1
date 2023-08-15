/*
Реализовать собственную функцию sleep.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	duration := time.Duration(2) * time.Second
	test(1, duration)
	test(2, duration)
	test(3, duration)
	test(4, duration)
}

func test(i int, d time.Duration) {
	fmt.Printf("start solution%d: %s\n", i, time.Now())
	switch i {
	case 1:
		solution1(d)
	case 2:
		solution2(d)
	case 3:
		solution3(d)
	case 4:
		solution4(d)
	}
	fmt.Printf("end: %s\n\n", time.Now())
}

// Использование  time.After  для получения канала,
// который будет отправлять значение после истечения заданного времени
func solution1(d time.Duration) {
	<-time.After(d)
}

// Использование  time.NewTicker  для создания тикера,
// который будет отправлять сигналы с заданной периодичностью.
// Мы будем ожидать нужное количество сигналов
func solution2(d time.Duration) {
	seconds := int(d / time.Second)
	ticker := time.NewTicker(d)
	for i := 0; i < seconds; i++ {
		<-ticker.C
	}
	ticker.Stop()
}

// Использование  time.Now()  для получения текущего времени,
// а затем ожидание пока время не достигнет конечного значения
func solution3(d time.Duration) {
	end := time.Now().Add(d)
	for time.Now().Before(end) {
	}
}

// Использование time.NewTimer для создания таймера,
// который отправит в канал время по истечению времени
func solution4(d time.Duration) {
	done := make(chan bool)
	timer := time.NewTimer(d)
	go func() {
		<-timer.C
		done <- true
	}()

	<-done
}
