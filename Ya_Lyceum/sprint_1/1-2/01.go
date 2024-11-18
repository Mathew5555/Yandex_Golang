package main

import (
	"fmt"
)

type Book struct {
	Title  string
	Author string
	Year   int
	Genre  string
}

func NewBook(title, author string, year int, genre string) *Book {
	return &Book{Title: title, Author: author, Year: year, Genre: genre}
}

func main() {
	// Используем конструктор
	book := NewBook("Alice", "Tom", 1922, "comedy")

	fmt.Printf("Название: %s, Автор: %s, Year: %d, Genre: %s\n", book.Title, book.Author, book.Year, book.Genre)
}
