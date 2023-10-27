package main

import (
	"html/template"
	"log"
	"net/http"
)

type Film struct {
	Title    string
	Director string
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film{
			"Films": {
				{Title: "The Matrix", Director: "Wachowski"},
				{Title: "The Matrix Reloaded", Director: "Wachowski"},
				{Title: "The Matrix Revolutions", Director: "Wachowski"},
			},
		}
		tmpl.Execute(w, films)
	})

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")

		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
