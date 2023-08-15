package main

import (
	"fmt"
	"testing"
	"time"
)

var p1 = NewPoint(1<<32, 1<<32)
var p2 = NewPoint(1<<5, 1<<2)

func TestDistance1_Performance(t *testing.T) {
	start := time.Now()
	distance1(p1, p2)
	elapsed := time.Since(start)

	fmt.Printf("Время выполнения distance1: %s\n", elapsed)
}

func TestDistance2_Performance(t *testing.T) {
	start := time.Now()
	distance2(p1, p2)
	elapsed := time.Since(start)

	fmt.Printf("Время выполнения distance2: %s\n", elapsed)
}
