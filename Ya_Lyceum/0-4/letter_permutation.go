package main

import "fmt"

func main() {
	var num int
	fmt.Scanln(&num)
	summa := 1
	for i := range num {
		summa *= i + 1
	}
	fmt.Println(summa)
}
