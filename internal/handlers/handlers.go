package handlers

import (
	"bytes"
	"html/template"
	"net/http"
	"path/filepath"
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

func renderTemplate(w http.ResponseWriter, tmplName string) {
	templateCache, err := createTemplateCache()

	if err != nil {
		panic(err)
	}

	// templateCache["home.page.tmpl"]
	tmpl, ok := templateCache[tmplName+".page.tmpl"]

	if !ok {
		http.Error(w, "Le template n'existe pas !", http.StatusInternalServerError)
		return
	}

	buffer := new(bytes.Buffer)
	tmpl.Execute(buffer, nil)
	buffer.WriteTo(w)
}

func createTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/templates/*.page.tmpl")

	if err != nil {
		return cache, err
	}

	// fmt.Println(pages)

	for _, page := range pages {
		name := filepath.Base(page)
		tmpl := template.Must(template.ParseFiles(page))

		layouts, err := filepath.Glob("./ui/html/layouts/*.layout.tmpl")

		if err != nil {
			return cache, err
		}

		if len(layouts) > 0 {
			tmpl.ParseGlob("./ui/html/layouts/*.layout.tmpl")
		}

		cache[name] = tmpl
	}

	return cache, nil
}
