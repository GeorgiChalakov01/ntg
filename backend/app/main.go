package main

import (
	"os"
	"fmt"
	"net/http"
	"github.com/a-h/templ"
	"github.com/GeorgiChalakov01/ntg/backend/app/pages/home"
	"github.com/GeorgiChalakov01/ntg/backend/app/pages/documents"
)

func main() {
	http.Handle("/", http.RedirectHandler("/home", http.StatusSeeOther))
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request){
		templ.Handler(home.Home()).ServeHTTP(w, r)
	})
	http.HandleFunc("/documents", func(w http.ResponseWriter, r *http.Request){
		templ.Handler(documents.Documents()).ServeHTTP(w, r)
	})
	
	port := os.Getenv("BACKEND_PORT")
	if port == "" {
		port = "8080"
	}
	
	fmt.Printf("Serving application on port %s ...\n", port)
	http.ListenAndServe(":" + port, nil)
}
