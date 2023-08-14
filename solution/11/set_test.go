package main

import (
	"fmt"
	"testing"
	"time"
)

var sliceSize = 100000

func generateLargeData(size int) []any {
	data := make([]any, size)
	for i := 0; i < size; i++ {
		data[i] = i
	}
	return data
}

func TestSolution1_Performance(t *testing.T) {
	// Генерируем большие данные
	set1 := generateLargeData(sliceSize)
	set2 := generateLargeData(sliceSize)

	// Замеряем время выполнения функции
	start := time.Now()
	solution1(set1, set2)
	elapsed := time.Since(start)

	fmt.Printf("Время выполнения solution1: %s\n", elapsed)
}

func TestSolution2_Performance(t *testing.T) {
	// Генерируем большие данные
	set1 := generateLargeData(sliceSize)
	set2 := generateLargeData(sliceSize)

	// Замеряем время выполнения функции
	start := time.Now()
	solution2(set1, set2)
	elapsed := time.Since(start)

	fmt.Printf("Время выполнения solution2: %s\n", elapsed)
}
