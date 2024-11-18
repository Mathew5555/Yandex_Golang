package main

import (
	"fmt"
	"os"
	"strconv"
)

func run() error {
	args := os.Args
	for i := 1; i < 4; i++ {
		if _, err := strconv.Atoi(args[i]); err != nil {
			return err
			//fmt.Errorf("%s doesnt look like a number.\n", args[i])
		}
	}

	file, _ := os.Create("config.txt")
	defer file.Close()
	text := fmt.Sprintf("%sx%s %s%%", args[1], args[2], args[3])
	_, err := file.WriteString(text)
	if err != nil {
		return err
	}
	return nil
}
