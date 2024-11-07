package main

import "fmt"

func CalculateSeriesSum(n float64) float64 {
	if n == 1 {
		return 1
	}
	return 1/float64(n) + CalculateSeriesSum(n-1)
}

func main() {
	fmt.Println(CalculateSeriesSum(2))
}
