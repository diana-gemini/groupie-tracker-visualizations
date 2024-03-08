package groupie

import (
	"net/http"
	"text/template"
)

func MainPage(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		ErrorPage(w, r, http.StatusNotFound)
		return
	}

	templ, err := template.ParseFiles("ui/html/index.html", "ui/html/layout.html")
	if err != nil {
		ErrorPage(w, r, http.StatusNotFound)
		return
	}

	err = templ.Execute(w, Groups)
	if err != nil {
		ErrorPage(w, r, http.StatusInternalServerError)
		return
	}
}
