package main

import (
	"fmt"
)

type Book struct {
	title  string
	author string
	pages  int
}

type ValidationError struct {
	Field string
	Value any
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("invalid %s : %v", e.Field, e.Value)
}

func (b Book) Validate() error {
	if b.pages <= 0 {
		return &ValidationError{"pages", b.pages}
	}
	if b.title == "" {
		return &ValidationError{"title", b.title}
	}
	return nil
}

func (b Book) String() string {
	if err := b.Validate(); err != nil {
		return fmt.Sprintf("Book(invalid): %v", err)
	}
	return fmt.Sprintf("%s by %s, %d pages", b.title, b.author, b.pages)
}

func main() {
	b := Book{title: "", author: "Orwell", pages: 0}
	fmt.Println(b)
}
