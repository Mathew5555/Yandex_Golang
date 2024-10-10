package main

import "fmt"

type Logger interface {
	Log()
}

type LogLevel string

const (
	Error LogLevel = "ERROR"
	Info  LogLevel = "INFO"
)

type Log struct {
	Level LogLevel
}

func (l Log) Log(message string) {
	switch l.Level {
	case Error:
		fmt.Printf("ERROR: %s\n", message)
	case Info:
		fmt.Printf("INFO: %s\n", message)
	default:
		fmt.Println("Something went wrong")
	}
}
