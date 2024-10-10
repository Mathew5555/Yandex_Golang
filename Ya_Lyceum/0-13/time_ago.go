package main

import (
	"fmt"
	"math"
	"time"
)

func FormatedText(num int) string {
	if num == 1 {
		return ""
	}
	return "s"
}

func TimeAgo(pastTime time.Time) string {
	durationSeconds := time.Now().Sub(pastTime).Seconds()
	minutes, hours, days, months, years := CalculateTime(durationSeconds)
	if durationSeconds < 60 {
		text := FormatedText(int(durationSeconds))
		return fmt.Sprintf("%d second%s ago", int(durationSeconds), text)
	} else if minutes < 60 {
		text := FormatedText(minutes)
		return fmt.Sprintf("%d minute%s ago", minutes, text)
	} else if hours < 24 {
		text := FormatedText(hours)
		return fmt.Sprintf("%d hour%s ago", hours, text)
	} else if days < 32 {
		text := FormatedText(days)
		return fmt.Sprintf("%d day%s ago", days, text)
	} else if months < 12 {
		text := FormatedText(months)
		return fmt.Sprintf("%d month%s ago", months, text)
	} else {
		text := FormatedText(years)
		return fmt.Sprintf("%d year%s ago", years, text)
	}
}

func CalculateTime(seconds float64) (int, int, int, int, int) {
	minutes := math.Round(seconds / 60)
	hours := math.Round(seconds / 3600)
	days := math.Round(seconds / 86400)
	months := math.Round(seconds / 2629440)
	years := math.Round(seconds / 31553280)
	return int(minutes), int(hours), int(days), int(months), int(years)
}

func main() {
	pastTime := time.Now().Add(-2 * time.Hour)
	result := TimeAgo(pastTime)
	fmt.Println(result)
}
