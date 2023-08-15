package main

import (
	"fmt"
	"testing"
	"time"
)

var stingSize = 300000

func generateLargeString(length int) string {
	// Генерируем строку заданной длины с повторяющимися символами
	str := ""
	for i := 0; i < length; i++ {
		str += "a"
	}
	return str
}

func TestSolution1_Performance(t *testing.T) {
	// Генерируем большую строку
	input := generateLargeString(stingSize)

	// Замеряем время выполнения функции
	start := time.Now()
	solution1(input)
	elapsed := time.Since(start)

	fmt.Printf("Время выполнения solution1: %s\n", elapsed)
}

func TestSolution2_Performance(t *testing.T) {
	// Генерируем большую строку
	input := generateLargeString(stingSize)

	// Замеряем время выполнения функции
	start := time.Now()
	solution2(input)
	elapsed := time.Since(start)

	fmt.Printf("Время выполнения solution2: %s\n", elapsed)
}
