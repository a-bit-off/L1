/*
Реализовать пересечение двух неупорядоченных множеств.
*/
package main

import "fmt"

func main() {
	set1 := []any{1, 2, 2, 3, 4}
	set2 := []any{3, 4, 5, 6}

	fmt.Println(set1)
	fmt.Println(set2)
	fmt.Println(solution1(set1, set2))
	fmt.Println(solution2(set1, set2))

}

func solution1(set1, set2 []any) []any {
	size := (len(set1) + len(set2)) / 2
	set := make(map[any]struct{})
	intersection := make([]any, 0, size)

	// добоавляем в set все из первого множеста
	for _, v := range set1 {
		set[v] = struct{}{}
	}

	// добавляем все что совпадает в set и во стором множестве
	for _, v := range set2 {
		if _, ok := set[v]; ok {
			intersection = append(intersection, v)
		}
	}

	return intersection
}

// менее производительное решение
func solution2(set1, set2 []any) []any {
	size := (len(set1) + len(set2)) / 2
	intersection := make([]any, 0, size)

	for _, v1 := range set1 {
		for _, v2 := range set2 {
			if v1 == v2 {
				intersection = append(intersection, v1) // добавляем все что совпадает в двух множествах
				break
			}
		}

	}

	return intersection
}
