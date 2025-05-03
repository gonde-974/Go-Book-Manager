package models

import "strings"

type Book struct {
	ID int
	Title  string
	Author string
	Year   int
}

type BookManager struct {
	Books []Book
}

func (bm *BookManager) AddBook(title, author string, year int) {
	id := len(bm.Books) + 1
	bm.Books = append(bm.Books, Book{ID: id, Title: title, Author: author, Year: year})
}

func (bm *BookManager) DeleteBook(id int) {
	for i, b := range bm.Books {
		if b.ID == id {
			bm.Books = append(bm.Books[:i], bm.Books[i+1:]...)
			break
		}
	}
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
