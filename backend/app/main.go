package main

import (
	"fmt"
	"net/http"
	"os"

	aboutus "github.com/GeorgiChalakov01/ntg/backend/app/pages/about_us"
	"github.com/GeorgiChalakov01/ntg/backend/app/pages/documents"
	"github.com/GeorgiChalakov01/ntg/backend/app/pages/home"
	"github.com/a-h/templ"
)

func main() {
	http.Handle("/", http.RedirectHandler("/home", http.StatusSeeOther))
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(home.Home()).ServeHTTP(w, r)
	})
	// Add route for about us page
	http.HandleFunc("/about-us", func(w http.ResponseWriter, r *http.Request) {
		aboutus.AboutUs().Render(r.Context(), w)
	})
	http.HandleFunc("/documents", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(documents.Documents()).ServeHTTP(w, r)
	})

	port := os.Getenv("BACKEND_PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Serving application on port %s ...\n", port)
	http.ListenAndServe(":"+port, nil)
}
