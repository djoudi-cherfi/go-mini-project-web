package main

import (
	"fmt"
	"net/http"

	"github.com/djoudi-cherfi/go-mini-project-web/internal/handlers"
)

const port = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/contact", handlers.Contact)

	fs := http.FileServer(http.Dir("ui/static"))
	http.Handle("/ui/static/", http.StripPrefix("/ui/static", fs))

	fmt.Println("(http://localhost:8080) - Server running on port", port)
	http.ListenAndServe(port, nil)
}
