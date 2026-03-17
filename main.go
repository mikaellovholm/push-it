package main

import (
	"html/template"
	"math/rand"
	"net/http"
)

var names = []string{
	"Alice", "Bob", "Charlie", "Diana", "Eve",
	"Frank", "Grace", "Hank", "Iris", "Jack",
	"Karen", "Leo", "Mona", "Nate", "Olivia",
	"Pete", "Quinn", "Rosa", "Sam", "Tina",
}

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

func handler(w http.ResponseWriter, r *http.Request) {
	name := names[rand.Intn(len(names))]
	tmpl.Execute(w, name)
}

func main() {
	http.HandleFunc("/", handler)
	println("push-it running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
