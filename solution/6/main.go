/*
Реализовать все возможные способы остановки выполнения горутины.
*/

package main

import (
	"context"
	"sync"
	"sync/atomic"
	"fmt"
	"time"
	"runtime"
)

func main() {

	mainWG := sync.WaitGroup{}
	mainWG.Add(1)
	defer mainWG.Wait()
	go func() {
		defer mainWG.Done()
		time.Sleep(time.Second * 7)
	}()

	// Solution1
	go func() {
		ctx, cancel := context.WithCancel(context.Background())
		go Solution1(ctx)
		time.Sleep(time.Second * 1)
		cancel() // stop goroutine
	}()

	// Solution2
	go func() {
		stopCh := make(chan int)
		go Solution2(stopCh)
		time.Sleep(time.Second * 2)
		stopCh<-1 // stop goroutine
	}()

	// Solution3
	go func() {
		wg := sync.WaitGroup{}
		wg.Add(1)
		go Solution3(&wg) // stop goroutine
		wg.Wait()
	}()

	// Solution4
	go func() {
		var stopFlag int32 = 0
		go Solution4(&stopFlag)
		time.Sleep(time.Second * 4)
		stopFlag = 1  // stop goroutine
	}()

	// Solution5
	go Solution5()

}

func Solution1(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Горутина остановилась за счет context")
			return
		}
	}
}

func Solution2(stopCh chan int) {
	for {
		select {
		case <-stopCh:
			fmt.Println("Горутина остановилась за счет канала")
			return
		}
	}
}

func Solution3(wg *sync.WaitGroup) {
	now := time.Now()
	for {
		if time.Since(now) > time.Duration(3)*time.Second {
			wg.Done()
			fmt.Println("Горутина остановилась за счет WaitGroup")
			return
		}
	}
}

func Solution4(stopFlag *int32) {
	for {
		if atomic.LoadInt32(stopFlag) == 1 {
			fmt.Println("Горутина остановилась за счет переменной")
			return
		}
	}
}

func Solution5() {
	now := time.Now()
	for {
		if time.Since(now) > time.Duration(5)*time.Second {
			fmt.Println("Горутина остановилась за счет runtime.Goexit()")
			runtime.Goexit()
		}
	}
}