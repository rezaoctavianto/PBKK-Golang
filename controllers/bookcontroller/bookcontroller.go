package bookcontroller

import (
	"Authors/entities"
	"Authors/models/authormodel"
	"Authors/models/bookmodel"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	books := bookmodel.GetAll()
	data := map[string]any{
		"books": books,
	}
	temp, err := template.ParseFiles("views/book/view.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/book/create.html")
		if err != nil {
			panic(err)
		}

		authors := authormodel.GetAll()
		data := map[string]any{
			"authors": authors,
		}

		temp.Execute(w, data)
	}
	if r.Method == "POST" {
		var book entities.Book

		authorId, err := strconv.Atoi(r.FormValue("author_id"))
		if err != nil {
			panic(err)
		}

		book.Title = r.FormValue("title")
		book.Author.Id = int(authorId)
		book.Genre = r.FormValue("genre")
		book.Description = r.FormValue("description")
		book.Updated_At = time.Now()
		book.Added_At = time.Now()

		if ok := bookmodel.Create(book); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, "/books", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/book/edit.html")
		if err != nil {
			panic(err)
		}

		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		book := bookmodel.Detail(id)
		authors := authormodel.GetAll()
		data := map[string]any{
			"authors": authors,
			"book":    book,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var book entities.Book

		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		authorId, err := strconv.Atoi(r.FormValue("author_id"))
		if err != nil {
			panic(err)
		}

		book.Title = r.FormValue("title")
		book.Author.Id = int(authorId)
		book.Genre = r.FormValue("genre")
		book.Description = r.FormValue("description")
		book.Updated_At = time.Now()

		if ok := bookmodel.Update(id, book); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, "/books", http.StatusSeeOther)
	}
}

func Detail(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	book := bookmodel.Detail(id)
	data := map[string]any{
		"book": book,
	}

	temp, err := template.ParseFiles("views/book/detail.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	if err := bookmodel.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/books", http.StatusSeeOther)
}
