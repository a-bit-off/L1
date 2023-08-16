/*
Реализовать конкурентную запись данных в map.
*/

package main

import (
	"fmt"
	"sync"
)

type ConcurrencyMap struct {
	cMap       map[any]any
	sync.Mutex // используем RWMutex для предоставления горутине уникального доступа к участку памяти cMap
}

func main() {
	testMap := map[int]string{
		0: "Ноль",
		1: "Один",
		2: "Два",
		3: "Три",
		4: "Четыре",
		5: "Пять",
		6: "Шесть",
		7: "Семь",
		8: "Восемь",
		9: "Девять",
	}
	m := New()
	wg := sync.WaitGroup{}

	count := 0
	for k, v := range testMap {
		wg.Add(1)

		go func(k int, v string, count int) {
			defer wg.Done()
			m.Add(k, v) // ADD
			if count%2 == 0 {
				m.Update(k, v+"Update") // UPDATE
			} else if count%3 == 0 {
				m.Delete(k) // DELETE
			}
		}(k, v, count)

		count++
	}

	wg.Wait()

	for k, v := range m.cMap {
		fmt.Printf("k = %d, v = %s\n", k, v)
	}
}

func New() *ConcurrencyMap {
	return &ConcurrencyMap{cMap: make(map[any]any)}
}

func (cm *ConcurrencyMap) Add(key, value any) {
	cm.Lock()
	cm.cMap[key] = value
	cm.Unlock()
}

func (cm *ConcurrencyMap) Update(key, value any) {
	cm.Lock()
	if _, exists := cm.cMap[key]; exists {
		cm.cMap[key] = value
	}
	cm.Unlock()
}

func (cm *ConcurrencyMap) Delete(key any) {
	cm.Lock()
	delete(cm.cMap, key)
	cm.Unlock()
}
