package main

import "fmt"

type Book struct {
	title  string
	author string
	pages  int
}

func CompareBooks(b1 Book, b2 Book) Book {
	if b1.pages > b2.pages {
		return b1
	}
	return b2
}

func PrintBooks(books []Book) {
	for _, book := range books {
		fmt.Println(book.title)
	}
}

func main() {
	book1 := Book{title: "1984", pages: 300}
	book2 := Book{title: "Dune", pages: 600}

	longest := CompareBooks(book1, book2)
	fmt.Println(longest.title)

	books := []Book{
		{title: "1984", pages: 300},
		{title: "Dune", pages: 600},
		{title: "Foundation", pages: 400},
	}
	PrintBooks(books)
}
