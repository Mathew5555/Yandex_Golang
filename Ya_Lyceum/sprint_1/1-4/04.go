package main

import "os"

func ModifyFile(filename string, pos int, val string) {
	f, _ := os.OpenFile(filename, os.O_WRONLY, 0600)
	f.Seek(int64(pos), 0)
	f.WriteString(val)
}
