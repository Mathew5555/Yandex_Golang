package main

import (
	"bufio"
	"os"
)

func LineByNum(inputFilename string, lineNum int) string {
	f, _ := os.Open(inputFilename)
	fileScanner := bufio.NewScanner(f)
	var i = 0
	for fileScanner.Scan() {
		if lineNum != i {
			i += 1
			continue
		}
		line := fileScanner.Text()
		return line
	}
	return ""
}

func main() {}
