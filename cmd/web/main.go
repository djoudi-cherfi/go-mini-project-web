package main

import (
	"fmt"
	"net/http"

	"github.com/djoudi-cherfi/go-mini-project-web/config"
	"github.com/djoudi-cherfi/go-mini-project-web/internal/handlers"
)

func main() {
	var appConfig config.Config

	templateCache, err := handlers.CreateTemplateCache()

	if err != nil {
		panic(err)
	}

	appConfig.TemplateCache = templateCache
	appConfig.Port = ":8080"

	handlers.CreateTemplates(&appConfig)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/contact", handlers.Contact)

	fs := http.FileServer(http.Dir("ui/static"))
	http.Handle("/ui/static/", http.StripPrefix("/ui/static", fs))

	fmt.Println("(http://localhost:8080) - Server started on port", appConfig.Port)
	http.ListenAndServe(appConfig.Port, nil)
}
