package main

import "fmt"

type Describer interface {
	Describe() string
}

type Book struct {
	title string
	pages int
}

func (b Book) Describe() string {
	return fmt.Sprintf("Book: %s (%d pages)", b.title, b.pages)
}

type Movie struct {
	title    string
	director string
}

func (m Movie) Describe() string {
	return fmt.Sprintf("Movie: %s (dir: %s)", m.title, m.director)
}

type Song struct {
	artist   string
	duration int
}

func (s Song) Describe() string {
	return fmt.Sprintf("Song by %s (%d sec)", s.artist, s.duration)
}

func PrintDescription(d Describer) {
	fmt.Println(d.Describe())
}

func main() {
	PrintDescription(Book{title: "1984", pages: 300})
	PrintDescription(Movie{title: "Inception", director: "Nolan"})
	PrintDescription(Song{artist: "Radiohead", duration: 245})
}
