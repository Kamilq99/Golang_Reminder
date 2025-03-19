package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Year   int
}

type Library struct {
	Books []Book
}

func addBook(l *Library, book Book) {
	l.Books = append(l.Books, book)
}

func findBook(l *Library, author string) *Book {
	for _, book := range l.Books {
		if book.Author == author {
			return &book
		}
	}
	return nil
}

func main() {

	l := Library{
		Books: []Book{
			{Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Year: 1925},
			{Title: "To Kill a Mockingbird", Author: "Harper Lee", Year: 1960},
			{Title: "1984", Author: "George Orwell", Year: 1949},
		},
	}

	fmt.Println(l)

	newBook := Book{
		Title:  "The Catcher in the Rye",
		Author: "J.D. Salinger",
		Year:   1951,
	}

	addBook(&l, newBook)

	fmt.Println(l)

	foundBook := findBook(&l, "J.D. Salinger")

	if foundBook != nil {
		fmt.Printf("Found book: %s by %s (%d)\n", foundBook.Title, foundBook.Author, foundBook.Year)
	}
}
