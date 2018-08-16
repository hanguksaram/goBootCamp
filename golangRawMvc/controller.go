package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)


//ROUTES
const (
	index = "/"
	about = "/about"
	contact = "/contact"
)

type some struct {

}

func (s some) initRoutes (mux *httprouter.Router) {
	mux.GET(index, s.in)
	mux.GET(about, s.ab)
	mux.GET(contact, s.co)
}

//HANDLERS
func (s some) in(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func (s some) ab(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	tpl.ExecuteTemplate(w, "about.gohtml", nil)

}

func (s some) co(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	tpl.ExecuteTemplate(w, "contact.gohtml", nil)

}


