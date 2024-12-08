package homecontroller

import (
	"html/template"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, nil)
}