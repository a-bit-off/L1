/*
Реализовать бинарный поиск встроенными методами языка.
*/
package main

import (
	"errors"
	"fmt"
)

func main() {
	names := []string{
		"Alex", "Alice", "Beth", "Bob", "Caleb", "Charlie", "David",
		"Diana", "Emma", "Ethan", "Fiona", "Frank",
		"George", "Grace", "Hannah", "Henry", "Ian",
		"Isabella", "Jack", "Kate", "Liam", "Mia",
		"Noah", "Olivia", "Paul", "Quinn", "Ryan",
		"Sophia", "Thomas", "Uma", "Victor", "Wendy",
		"Xavier", "Yara", "Zoe",
	}

	find := "Zoe"
	index, err := binarySearch(names, find)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(index, names[index])
	}
}

func binarySearch(names []string, find string) (int, error) {
	low := 0
	high := len(names) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := names[mid]

		if guess == find {
			return mid, nil
		} else if guess > find {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return 0, errors.New("Не удалось найти!")
}
