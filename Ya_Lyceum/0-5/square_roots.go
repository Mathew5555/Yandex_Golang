package main

import (
	"fmt"
	"math"
)

func SqRoots() {
	var a, b, c float64
	fmt.Scanln(&a, &b, &c)
	discriminant := math.Sqrt(math.Pow(b, 2) - 4*a*c)
	if math.IsNaN(discriminant) || discriminant < 0 {
		fmt.Println(0, 0)
	} else if discriminant == 0 {
		fmt.Println(-b / (2 * a))
	} else {
		fmt.Println((-b-discriminant)/(2*a), (-b+discriminant)/(2*a))
	}
}

func main() {
	SqRoots()
}
