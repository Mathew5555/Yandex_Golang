package main

import "fmt"

func main() {
	var score int
	fmt.Scanln(&score)
	switch {
	case score > 0:
		fmt.Println("Число положительное")
	case score < 0:
		fmt.Println("Число отрицательное")
	default:
		fmt.Println("Введен ноль")
	}
}
