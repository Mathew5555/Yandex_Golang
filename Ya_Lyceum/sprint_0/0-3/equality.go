package main

import "fmt"

func main() {
	var a, b, c int
	_, err := fmt.Scanln(&a, &b, &c)
	if err != nil {
		fmt.Println("Некорректный ввод")
	} else if a == b && b == c {
		fmt.Println("Все числа равны")
	} else if a == b || b == c || c == a {
		fmt.Println("Два числа равны")
	} else {
		fmt.Println("Все числа разные")
	}
}
