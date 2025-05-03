package models

import "strings"

type Book struct {
	Title  string
	Author string
	Year   int
}

type BookManager struct {
	Books []Book
}

func (bm *BookManager) AddBook(title, author string, year int) {
	bm.Books = append(bm.Books, Book{title, author, year})
}

func (bm *BookManager) ListBooks() []Book {
	return bm.Books
}

func (bm *BookManager) SearchByAuthor(author string) []Book {
	var result []Book
	for _, b := range bm.Books {
		if strings.EqualFold(b.Author, author) {
			result = append(result, b)
		}
	}
	return result
}
