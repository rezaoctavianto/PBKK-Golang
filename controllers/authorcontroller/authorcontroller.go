package authorcontroller

import (
	"Authors/entities"
	"Authors/models/authormodel"
	"net/http"
	"strconv"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	authors := authormodel.GetAll()
	data := map[string]any{
		"authors": authors,
	}
	temp, err := template.ParseFiles("views/author/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, data)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/author/edit.html")
		if err != nil {
			panic(err)
		}

		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		
		if err != nil {
			panic(err)
		}

		author := authormodel.Detail(id)
		data := map[string]any{
			"author": author,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var author entities.Author

		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}
		author.Name = r.FormValue("name")
		author.DoB = r.FormValue("DoB")

		if ok := authormodel.Update(id, author); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/authors", http.StatusSeeOther)
	}
}
func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/author/create.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, nil)
	}
	if r.Method == "POST" {
		var author entities.Author

		author.Name = r.FormValue("name")
		author.DoB = r.FormValue("DoB")

		if ok := authormodel.Create(author); !ok {
			temp, _ := template.ParseFiles("views/author/create.html")
			temp.Execute(w, nil)
		}

		http.Redirect(w, r, "/authors", http.StatusSeeOther)
	}
}
func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	if err := authormodel.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/authors", http.StatusSeeOther)
}
