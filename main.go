package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

type user int

func init() {
	tpl = template.Must(template.ParseFiles("one.html"))
}

func (u user) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(w, "one.html", r.Form)
}

func main() {

	var d user
	http.ListenAndServe(":8080", d)
}
