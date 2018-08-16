package main

import (
		"github.com/julienschmidt/httprouter"
	"html/template"
		"net/http"
)

var tpl *template.Template
var s some
func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	mux := httprouter.New()
	s.initRoutes(mux)
	http.ListenAndServe(":8080", mux)
}

