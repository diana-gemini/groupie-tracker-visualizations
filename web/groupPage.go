package groupie

import (
	"log"
	"net/http"
	"strconv"
	"text/template"
)

type GroupPage struct {
	G               Groupie
	NextGroupie     int
	PreviousGroupie int
}

func Group(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/group" {
		ErrorPage(w, r, http.StatusNotFound)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		ErrorPage(w, r, http.StatusNotFound)
		return
	}
	if id > len(Groups) || id < 1 {
		ErrorPage(w, r, http.StatusNotFound)
		return
	}

	templ, err := template.ParseFiles("ui/html/groupPage.html", "ui/html/layout.html")
	if err != nil {
		log.Println(err)
		ErrorPage(w, r, http.StatusNotFound)
		return
	}

	err = templ.Execute(w, Groups[id-1])
	if err != nil {
		ErrorPage(w, r, http.StatusInternalServerError)
		return
	}
}
