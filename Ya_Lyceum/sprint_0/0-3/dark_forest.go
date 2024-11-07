package main

import "fmt"

func main() {
	var num int
	fmt.Scanln(&num)
	if num > 0 && num%2 == 1 {
		fmt.Println("Число положительное и нечетное")
	} else if num > 0 {
		fmt.Println("Число положительное и четное")
	} else if num < 0 && num%2 == -1 {
		fmt.Println("Число отрицательное и нечетное")
	} else if num < 0 {
		fmt.Println("Число отрицательное и четное")
	} else {
		fmt.Println("Число равно нулю")
	}

}
