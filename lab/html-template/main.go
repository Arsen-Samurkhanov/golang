package main

import (
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	Title   string
	Message string
	Items   []string
}

func main() {

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			Title:   "Go HTML Template",
			Message: "Hello from Golang 👋",
			Items:   []string{"Apple", "Banana", "Orange"},
		}

		tmpl.Execute(w, data)

	})

	log.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
