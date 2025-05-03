package models

import "strings"

type Book struct {
	ID     int
	Title  string
	Author string
	Year   int
}

type BookManager struct {
	Books []Book
	nextID int
}

// Конструктор: иницијализира нов менаџер со некои книги
func NewBookManager() *BookManager {
	return &BookManager{
		Books: []Book{
			{ID: 1, Title: "Na Drini ćuprija", Author: "Ivo Andric", Year: 1945},
			{ID: 2, Title: "Seobe", Author: "Miloš Crnjanski", Year: 1929},
			{ID: 3, Title: "Travnička hronika", Author: "Ivo Andric", Year: 1945},
		},
		nextID: 4, // следниот ID
	}
}

func (bm *BookManager) AddBook(title, author string, year int) {
	book := Book{ID: bm.nextID, Title: title, Author: author, Year: year}
	bm.Books = append(bm.Books, book)
	bm.nextID++
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
