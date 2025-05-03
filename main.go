package main

import (
	"go-book-manager/routes"
	"net/http"
)

func main() {
	routes.RegisterRoutes()
	http.ListenAndServe(":8080", nil)
}
