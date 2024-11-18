package main

import "fmt"

type User2 struct {
	Name   string
	Age    int
	Active bool
}

func NewUser(name string) *User2 {
	return &User2{Name: name, Age: 18, Active: true}
}

func main() {
	user := NewUser("Matvey")
	fmt.Printf("Name: %s", user.Name)
}
