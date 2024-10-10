package main

import (
	"fmt"
	"time"
)

func NextWorkday(start time.Time) time.Time {
	start = start.Add(time.Hour * 24)
	for start.Weekday() > 5 || start.Weekday() == 0 {
		start = start.Add(time.Hour * 24)
	}
	return start
}

func main() {
	start := time.Date(2023, time.October, 6, 0, 0, 0, 0, time.UTC) // A Saturday
	nextWorkday := NextWorkday(start)
	fmt.Println(nextWorkday)
}
