package main

import "fmt"

type Book struct {
	title  string
	author string
	pages  int
}

func (b Book) Info() string {
	return fmt.Sprintf("%s by %s (%d pages)", b.title, b.author, b.pages)
}

func (b *Book) AddPages(n int) {
	b.pages += n
}

func main() {
	book := Book{title: "The Great Gatsby", author: "F. Scott Fitzgerald", pages: 180}
	fmt.Println(book.Info())

	book.AddPages(50)
	fmt.Println(book.pages)
}
