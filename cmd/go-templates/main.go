package main

import (
	"html/template"
	"net/http"
)

var tmpl *template.Template

type Todo struct {
	Item string
	Done bool
}

type PageData struct {
	Title string
	Todos []Todo
}

func todo(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "TODO List",
		Todos: []Todo{
			{Item: "Install Go", Done: true},
			{Item: "Learn Go", Done: false},
			{Item: "Make an API with Go", Done: false},
			{Item: "Make a Site with Go Templates", Done: false},
		},
	}
}

func main() {
	mux := http.NewServeMux()
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	mux.HandleFunc("/todo", todo)
}
