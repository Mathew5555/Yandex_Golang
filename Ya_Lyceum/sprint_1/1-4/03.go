package main

import (
	"fmt"
	"io"
	"os"
)

func CopyFilePart(inputFileName, outFileName string, startPos int) error {
	f, err := os.Open(inputFileName)
	defer f.Close()
	if err != nil {
		return err
	}
	f2, err2 := os.Create(outFileName)
	if err2 != nil {
		return err2
	}
	f.Seek(int64(startPos), 0)
	buffer := make([]byte, 1024)
	for {
		n, err := f.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		f2.Write(buffer[:n])
	}
	err = f2.Close()
	return err
}

func main() {
	err := CopyFilePart("C:\\Users\\Матвей\\GolandProjects\\Yandex_Golang\\Ya_Lyceum\\sprint_1\\1-4\\file.txt", "", 3)
	fmt.Println(err)
}
