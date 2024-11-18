package main

import (
	"fmt"
	"io"
	"os"
)

func ReadContent(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil && err != io.EOF {
		return ""
	}
	return string(data)
}

func main() {
	//ReadContent("file.txt")
	fmt.Println(ReadContent("C:\\Users\\Матвей\\GolandProjects\\Yandex_Golang\\Ya_Lyceum\\sprint_1\\1-4\\file.txt"))
}
