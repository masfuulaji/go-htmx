package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	fmt.Println("Hello, World!")

	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film{
			"films": {
				{Title: "The Matrix", Director: "Wachowski"},
				{Title: "The Matrix Reloaded", Director: "Wachowski"},
				{Title: "The Matrix Revolutions", Director: "Wachowski"},
			},
		}
		tmpl.Execute(w, films)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
