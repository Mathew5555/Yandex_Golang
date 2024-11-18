package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"time"
)

var (
	EmptyResult = errors.New("empty result")
)

func ExtractLog(inputFileName string, start, end time.Time) ([]string, error) {
	var ans []string
	f, _ := os.Open(inputFileName)
	defer f.Close()
	fileScanner := bufio.NewScanner(f)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		fmt.Println(line[:10])
		dateLine, _ := time.Parse("02.01.2006", line[:10])
		if dateLine.Before(start) {
			continue
		} else if dateLine.After(end) {
			break
		}
		ans = append(ans, line)
	}
	if len(ans) == 0 {
		return ans, EmptyResult
	}
	return ans, nil
}

func main() {
	_, err := ExtractLog("C:\\Users\\Матвей\\GolandProjects\\Yandex_Golang\\Ya_Lyceum\\sprint_1\\1-4\\file.txt", time.Date(2022, 12, 14, 0, 0, 0, 0, time.UTC), time.Date(2022, 12, 15, 0, 0, 0, 0, time.UTC))
	fmt.Println(err)
}
