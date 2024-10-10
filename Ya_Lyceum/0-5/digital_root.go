package main

import "fmt"

func SumDigitsRecursive2(n int) int {
	if n == 0 {
		return 0
	}
	return n%10 + SumDigitsRecursive2(n/10)
}

func CalculateDigitalRoot(n int) int {
	if n < 10 {
		return n
	}
	return CalculateDigitalRoot(SumDigitsRecursive2(n))
}

func main() {
	fmt.Println(CalculateDigitalRoot(12345))
}
