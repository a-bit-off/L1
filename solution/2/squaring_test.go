package main

import (
	"fmt"
	"testing"
	"time"
)

var size = 10

func Test_Squaring_Solution_1(t *testing.T) {
	numbers := make([]int, size)
	for i := range numbers {
		numbers[i] = i
	}

	start := time.Now()
	Squaring_Solution_1(numbers)
	elapsed := time.Since(start)

	fmt.Printf("Время выполнения solution1: %s\n", elapsed)
}

func Test_Squaring_Solution_2(t *testing.T) {
	numbers := make([]int, size)
	for i := range numbers {
		numbers[i] = i
	}

	start := time.Now()
	Squaring_Solution_2(numbers)
	elapsed := time.Since(start)

	fmt.Printf("Время выполнения solution2: %s\n", elapsed)
}

func Test_Squaring_Solution_3(t *testing.T) {
	numbers := make([]int, size)
	for i := range numbers {
		numbers[i] = i
	}

	start := time.Now()
	Squaring_Solution_3(numbers)
	elapsed := time.Since(start)

	fmt.Printf("Время выполнения solution3: %s\n", elapsed)
}

func Test_Squaring_Solution_4(t *testing.T) {
	numbers := make([]int, size)
	for i := range numbers {
		numbers[i] = i
	}

	start := time.Now()
	Squaring_Solution_4(numbers)
	elapsed := time.Since(start)

	fmt.Printf("Время выполнения solution4: %s\n", elapsed)
}
