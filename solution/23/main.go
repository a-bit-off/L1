/*
Удалить i-ый элемент из слайса.
*/
package main

import (
	"fmt"
	"log"
)

func main() {
	slice := []int{1, 2, 3, 4, 5}
	res, err := delete(slice, 4)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res)
}

func delete(slice []int, i int) ([]int, error) {
	if i < 0 || i > len(slice)-1 {
		return nil, fmt.Errorf("Index out of range")
	}
	// объединяем два срезу исключая позицию i
	return append(slice[:i], slice[i+1:]...), nil
}
