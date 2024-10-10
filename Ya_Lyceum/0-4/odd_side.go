package main

import "fmt"

func main() {
	var num int
	fmt.Scanln(&num)
	summa := 0
	for i := range num {
		if (i+1)&1 == 1 {
			summa += i + 1
		}
	}
	if num < 0 {
		fmt.Println("Некорректный ввод")
	} else {
		fmt.Println(summa)
	}
}
