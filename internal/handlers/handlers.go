package handlers

import (
	"bytes"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/djoudi-cherfi/go-mini-project-web/config"
	"github.com/djoudi-cherfi/go-mini-project-web/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	names := make(map[string]string)
	names["owner"] = "John"

	renderTemplate(w, "home", &models.TemplateData{
		StringData: names,
	})
}

func About(w http.ResponseWriter, r *http.Request) {
	age := make(map[string]int)
	age["owner"] = 30

	renderTemplate(w, "about", &models.TemplateData{
		IntData: age,
	})
}

func Contact(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "contact", &models.TemplateData{})
}

var appConfig *config.Config

func CreateTemplates(app *config.Config) {
	appConfig = app
}

func renderTemplate(w http.ResponseWriter, tmplName string, td *models.TemplateData) {
	templateCache := appConfig.TemplateCache

	// templateCache["home.page.tmpl"]
	tmpl, ok := templateCache[tmplName+".page.tmpl"]

	if !ok {
		http.Error(w, "Le template n'existe pas !", http.StatusInternalServerError)
		return
	}

	buffer := new(bytes.Buffer)
	tmpl.Execute(buffer, td)
	buffer.WriteTo(w)
}

func CreateTemplateCache() (map[string]*template.Template, error) {
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
