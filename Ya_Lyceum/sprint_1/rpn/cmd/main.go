package main

import (
	"errors"
	"github.com/Mathew5555/Yandex_Golang/Ya_Lyceum/sprint_1/rpn/internal/application"
)

var (
	NotPositive = errors.New("Numbers must be more than 0")
)

type World struct {
	Height int
	Width  int
	Cells  [][]bool
}

func NewWorld(height, width int) (*World, error) {
	if height < 0 || width < 0 {
		return nil, NotPositive
	}
	arr := make([][]bool, height)
	for i := 0; i < height; i++ {
		arr[i] = make([]bool, width)
	}
	myWorld := World{Height: height, Width: width, Cells: arr}
	return &myWorld, nil
}

func main() {
	app := application.New()
	err := app.Run()
	if err != nil {
		return
	}
}
