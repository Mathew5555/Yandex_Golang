package main

import "fmt"

func main() {
	var stars int
	fmt.Scanln(&stars)
	switch {
	case stars >= 1000:
		fmt.Println("Число больше или равно 1000")
	case 100 <= stars && stars < 1000:
		fmt.Println("Число меньше 1000")
	case 10 <= stars && stars < 100:
		fmt.Println("Число меньше 100")
	case 0 <= stars && stars < 10:
		fmt.Println("Число меньше 10")
	default:
		fmt.Println("Некорректный ввод")
	}
}
