package main

func SumOfArray(array [6]int) int {
	summa := 0
	for _, el := range array {
		summa += el
	}
	return summa
}
