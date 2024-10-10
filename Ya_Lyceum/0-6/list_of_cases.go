package main

import "fmt"

func PrettyArrayOutput(array [9]string) {
	for i, el := range array {
		if i < 7 {
			fmt.Println(i+1, "я уже сделал:", el)
		} else {
			fmt.Println(i+1, "не успел сделать:", el)
		}
	}
}
