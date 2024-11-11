package main

import (
	"fmt"
	"net/http"
	"os"

	x "jemi/Functions"
)

func main() {
	fmt.Println("server started at http://localhost:8081")
	http.HandleFunc("/", x.HomeHandler)
	http.HandleFunc("/Artist", x.Second_Page_Handler)
	http.HandleFunc("/Search_Bar",x.Search_Bar)
	http.HandleFunc("/Filter",x.Filter)
	http.HandleFunc("/static/", Static)
	http.ListenAndServe(":8081", nil)
}

func Static(w http.ResponseWriter, r *http.Request) {
	fs := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	_, err := os.Stat("." + r.URL.Path)
	if err != nil {
		x.ErrorHandler(w, r, 404)
		return
	}
	if r.URL.Path == "/static/" {
		x.ErrorHandler(w, r, 403)
		return
	}
	fs.ServeHTTP(w, r)
}
