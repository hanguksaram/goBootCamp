package main

import (
	"net/http"

	"github.com/GoesToEleven/udemyTraining/03/controllers"
	"github.com/GoesToEleven/udemyTraining/03/lib"
	"github.com/GoesToEleven/udemyTraining/03/middleware"
)

type Startup struct{}

func (st *Startup) Initialize() {
	mainConroller := &controllers.MainController{}
	handlers := mainConroller.InitController()
	middleware := lib.CreateNewMiddleware(middleware.CustomMiddleWare1)
	middlewaredHandlers := lib.ApplyMiddleware(handlers, middleware)
	for _, hn := range middlewaredHandlers {
		http.Handle(hn.Route, hn.Handler)
	}
}
