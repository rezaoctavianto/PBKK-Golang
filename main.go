package main

import (
	"Authors/config"
	"Authors/controllers/authorcontroller"
	"Authors/controllers/bookcontroller"
	"Authors/controllers/homecontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	//1.Home
	http.HandleFunc("/", homecontroller.Welcome)

	//2.Books
	http.HandleFunc("/books", bookcontroller.Index)
	http.HandleFunc("/books/add", bookcontroller.Add)
	http.HandleFunc("/books/edit", bookcontroller.Edit)
	http.HandleFunc("/books/delete", bookcontroller.Delete)

	//3.Authors
	http.HandleFunc("/authors", authorcontroller.Index)
	http.HandleFunc("/authors/add", authorcontroller.Add)
	http.HandleFunc("/authors/edit", authorcontroller.Edit)
	http.HandleFunc("/authors/delete", authorcontroller.Delete)

	log.Println("Server running on port :8000")
	http.ListenAndServe(":8000", nil)
}