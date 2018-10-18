package main

import (
	"net/http"

	"github.com/GoesToEleven/udemyTraining/03/controllers"
	"github.com/GoesToEleven/udemyTraining/03/lib"
	"github.com/GoesToEleven/udemyTraining/03/middleware"
)

type Startup struct{}

func (st *Startup) Initialize() {
	//factories
	mwFactory := &middleware.MiddlewareFactory{}
	mainConroller := &controllers.MainController{}

	//preparing middlewares
	mwCors := mwFactory.PrepareMiddleware(middleware.AllowCors)

	//controllers initializing
	handlers := mainConroller.InitController()

	//create httprequest handling chaing
	middlewaredHandlers := lib.ApplyMiddleware(handlers, mwCors)

	//init handlers
	for _, hn := range middlewaredHandlers {
		http.Handle(hn.Route, hn.Handler)
	}
}
