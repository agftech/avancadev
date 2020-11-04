package main

import (
	//"fmt"
	"html/template"
	"net/http"
)

type Result struct {
	Status string
}

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":9090", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "<h1>Hello World</h1>")
	t := template.Must(template.ParseFiles("templates/home.html"))
	t.Execute(w, "")
}