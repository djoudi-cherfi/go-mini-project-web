package handlers

import (
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about")
}

func Contact(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "contact")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("./ui/html/templates/" + tmpl + ".page.tmpl")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	t.Execute(w, nil)
}
