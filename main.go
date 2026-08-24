package main

import "fmt"

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

func (b Book) Validate() error 	{
	
}

func main() {
	b := Book{title: "", author: "Orwell", pages: 300}
	fmt.Println(b)
}