package main

import (
	"fmt"
	"math"
)

func FindMaxKey(m map[int]int) int {
	var res int = math.MinInt
	for k, _ := range m {
		res = max(res, k)
	}
	return res
}

func main() {
	m := map[int]int{
		-7: 1,
		-1: 38,
	}
	fmt.Println(FindMaxKey(m))
}
