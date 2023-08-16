/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
*/
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	input := make(chan int)
	n := CountOfWorkers()

	// запускаем горутину которая будет слушать основной поток (ловим Ctrl+C)
	go ListenerForSignals(cancel)

	for i := 0; i < n; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			Worker(ctx, input)
		}()
	}

	go Writer(ctx, input)

	wg.Wait()
}

func CountOfWorkers() int {
	count := flag.Int("n", 10, "count of workers")
	flag.Parse()
	return *count
}

func ListenerForSignals(cancel context.CancelFunc) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh  // ждем записи в канал
	cancel() // завершаем программу
}

func Worker(ctx context.Context, input chan int) {
	for {
		select {
		case <-ctx.Done(): // в случае получения сигнала безопасно завершаем горутины
			return
		default: // считываем с канала данные
			if val, ok := <-input; ok {
				fmt.Println(val)
			}
		}
	}
}

func Writer(ctx context.Context, input chan int) {
	defer close(input)
	counter := 0

	for {
		select {
		case <-ctx.Done(): // в случае получения сигнала безопасно завершаем горутины
			return
		default: // иначе пишем в канал
			input <- counter
			counter++
			time.Sleep(time.Second * 1)
		}
	}
}
