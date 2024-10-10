package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanln(&a, &b)
	if a > 0 && b > 0 {
		fmt.Println("Оба числа положительные")
	} else if a > 0 && b < 0 || a < 0 && b > 0 {
		fmt.Println("Одно число положительное, а другое отрицательное")
	} else if a == 0 || b == 0 {
		fmt.Println("Одно из чисел равно нулю")
	} else {
		fmt.Println("Оба числа отрицательные")
	}
}
