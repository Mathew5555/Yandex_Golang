package main

import "fmt"

func main() {
	var num int
	fmt.Scanln(&num)
	var summa int
	for i := range num {
		if (i+1)%3 != 0 && (i+1)%5 != 0 {
			summa += i + 1
		}
	}
	fmt.Println(summa)
}
