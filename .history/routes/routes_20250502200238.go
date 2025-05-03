package routes

import (
	"html/template"
	"net/http"
	"strconv"

	"go-book-manager/models"
)

var manager = models.BookManager{
	Books: []models.Book{
		{"Na Drini ćuprija", "Ivo Andric", 1945},
		{"Seobe", "Miloš Crnjanski", 1929},
		{"Travnička hronika", "Ivo Andric", 1945},
	},
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func RegisterRoutes() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/search", searchHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", manager.ListBooks())
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		title := r.FormValue("title")
		author := r.FormValue("author")
		year, _ := strconv.Atoi(r.FormValue("year"))
		manager.AddBook(title, author, year)
		templates.ExecuteTemplate(w, "list.html", manager.ListBooks())
	}
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	author := r.FormValue("author")
	books := manager.SearchByAuthor(author)
	templates.ExecuteTemplate(w, "list.html", books)
}
