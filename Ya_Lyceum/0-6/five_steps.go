package main

import "fmt"

func FiveSteps(array [5]int) []int {
	reversed := make([]int, len(array))
	for i := 0; i < len(array); i++ {
		reversed[i] = array[len(array)-i-1]
	}
	return reversed
}

func main() {
	input := [5]int{1, 2, 3, 4, 5}
	output := FiveSteps(input)
	for i := 0; i < len(output); i++ {
		fmt.Printf("%d ", output[i])
	}
}
