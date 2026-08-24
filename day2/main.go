package main

import "fmt"

type Book struct {
	title  string
	author string
	pages  int
	rating int
}

func (b Book) ShortDescription() string {
	title := b.title
	if len(title) > 50 {
		title = title[:50]
	}
	return title + " — " + b.author
}

func (b Book) IsLong() bool {
	return b.pages > 300
}

func (b Book) Summarize() string {
	return fmt.Sprintf("%s (%d pages)", b.title, b.pages)
}

func (b *Book) Rate(stars int) {
	b.rating = stars
}

func main() {
	book := Book{title: "Dune", author: "Herbert", pages: 600}

	fmt.Println(book.ShortDescription())
	fmt.Println(book.IsLong())
	fmt.Println(book.Summarize())

	book.Rate(5)
	fmt.Println(book.rating)
}
