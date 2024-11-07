package main

import (
	"fmt"
)

func CountingSort(contacts []string) map[string]int {
	res := make(map[string]int)
	for _, el := range contacts {
		res[el]++
	}
	return res
}

func main() {
	ls := []string{"a", "a", "b"}
	fmt.Println(CountingSort(ls))
}
