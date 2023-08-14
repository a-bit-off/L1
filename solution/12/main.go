/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/
package main

import (
	"12/set"
	"fmt"
)

func main() {
	set := set.New()

	set.Add("cat", "cat", "dog", "cat", "tree")
	set.Delete("tree")

	fmt.Println(set.Contain("dog"))
	fmt.Println(set.GetKeys())
}
