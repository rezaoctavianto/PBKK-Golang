package authorcontroller

import (
	"Authors/models/authormodel"
	"net/http"
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

}
func Add(w http.ResponseWriter, r *http.Request) {

}
func Delete(w http.ResponseWriter, r *http.Request) {

}
