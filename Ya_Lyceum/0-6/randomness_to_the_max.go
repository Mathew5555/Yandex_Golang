package main

func FindMinMaxInArray(array [10]int) (int, int) {
	min, max := array[0], array[0]
	for _, el := range array {
		if el < min {
			min = el
		}
		if el > max {
			max = el
		}
	}
	return min, max
}
