package main

func SumOfValuesInMap(m map[int]int) int {
	var sum int = 0
	for _, val := range m {
		sum += val
	}
	return sum
}
