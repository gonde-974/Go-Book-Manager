package routes

import (
	"go-book-manager/models"
	"html/template"
	"net/http"
	"strconv"
)

var templates = template.Must(template.ParseFiles(
	"templates/index.html",
	"templates/list.html",
))

var manager = models.BookManager()

// Home Page - го прикажува целиот интерфејс
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", manager.ListBooks())
}

// Додавање на нова книга преку форма (HTMX)
func AddHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		title := r.FormValue("title")
		author := r.FormValue("author")
		yearStr := r.FormValue("year")
		year, err := strconv.Atoi(yearStr)
		if err == nil {
			manager.AddBook(title, author, year)
		}
		templates.ExecuteTemplate(w, "list.html", manager.ListBooks())
	}
}

// Пребарување по автор (HTMX)
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	author := r.URL.Query().Get("author")
	results := manager.SearchByAuthor(author)
	templates.ExecuteTemplate(w, "list.html", results)
}

// Бришење на книга по ID (HTMX)
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		idStr := r.FormValue("id")
		id, err := strconv.Atoi(idStr)
		if err == nil {
			manager.DeleteBook(id)
		}
		templates.ExecuteTemplate(w, "list.html", manager.ListBooks())
	}
}

// RegisterRoutes ја повикуваш од main.go
func RegisterRoutes() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/add", AddHandler)
	http.HandleFunc("/search", SearchHandler)
	http.HandleFunc("/delete", DeleteHandler)

	// За статички фајлови (CSS, JS)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
}
