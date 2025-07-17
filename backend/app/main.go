package main


import (
	"os"
	"fmt"
	"net/http"
	"github.com/a-h/templ"
	//"github.com/jackc/pgx/v5"
	"github.com/GeorgiChalakov01/ntg/backend/app/pages/home"
)

func main() {
	http.Handle("/", http.RedirectHandler("/home", http.StatusSeeOther))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request){
		templ.Handler(home.Home()).ServeHTTP(w, r)
	})

	port := os.Getenv("BACKEND_PORT")
	fmt.Printf("Serving application on port %s ...\n", port)
	http.ListenAndServe(":" + port, nil)
}
