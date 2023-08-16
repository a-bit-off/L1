/*
Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

*/

package main

import (
	"context"
	"flag"
	"fmt"
	"sync"
	"time"
)

func main() {
	n := Timeout()
	// создаем канал который отправит сигнал о завершении через n секунд
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(n)*time.Second)
	defer cancel()
	input := make(chan int)
	wg := sync.WaitGroup{}

	go Write(ctx, input)

	wg.Add(1)
	go func() {
		defer wg.Done()
		Read(ctx, input)
	}()

	wg.Wait()
}

func Timeout() int {
	n := flag.Int("n", 5, "timeout")
	flag.Parse()
	return *n
}

func Read(ctx context.Context, inout chan int) {
	for {
		select {
		case <-ctx.Done(): // безопасное завершение горутины
			return
		default:
			if val, ok := <-inout; ok {
				fmt.Println(val)
			}
		}
	}
}

func Write(ctx context.Context, input chan int) {
	defer close(input)
	counter := 0

	for {
		select {
		case <-ctx.Done(): // безопасное завершение горутины
			return
		default:
			input <- counter
			counter++
			time.Sleep(time.Second * 1)
		}
	}
}
