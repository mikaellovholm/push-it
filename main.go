package main

import (
	_ "embed"
	"html/template"
	"math/rand"
	"net/http"
)

//go:embed templates/index.html
var indexHTML string

var names = []string{
	"Alice", "Bob", "Charlie", "Diana", "Eve",
	"Frank", "Grace", "Hank", "Iris", "Jack",
	"Karen", "Leo", "Mona", "Nate", "Olivia",
	"Pete", "Quinn", "Rosa", "Sam", "Tina",
}

var tmpl = template.Must(template.New("index").Parse(indexHTML))

func handler(w http.ResponseWriter, r *http.Request) {
	name := names[rand.Intn(len(names))]
	tmpl.Execute(w, name)
}

func main() {
	http.HandleFunc("/", handler)
	println("push-it running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
