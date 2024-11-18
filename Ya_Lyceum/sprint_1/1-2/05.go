package main

import "fmt"

var (
	cnt = 0
)

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

type Report struct {
	User
	ReportID int
	Date     string
}

func CreateReport(user User, reportDate string) Report {
	cnt++
	return Report{User: user, Date: reportDate, ReportID: cnt}
}

func PrintReport(report Report) {
	fmt.Printf(":D")
}

func GenerateUserReports(users []User, reportDate string) []Report {
	var reports []Report
	for _, user := range users {
		reports = append(reports, CreateReport(user, reportDate))
	}
	return reports
}
