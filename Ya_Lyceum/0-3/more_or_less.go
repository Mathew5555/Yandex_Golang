package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanln(&a, &b)
	switch {
	case a > b:
		fmt.Println("Первое число больше второго")
	case a < b:
		fmt.Println("Второе число больше первого")
	default:
		fmt.Println("Числа равны")
	}
}
