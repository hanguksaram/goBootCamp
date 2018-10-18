package controllers

import (
	"fmt"
	"net/http"

	"github.com/GoesToEleven/udemyTraining/03/lib"
)

const (
	indexRoute = "/"
)

type MainController struct{}

func (c *MainController) InitController() []*lib.RequestHandler {
	indexHandler := &lib.RequestHandler{Route: indexRoute, Handler: indexHandler, IsMiddlewared: true}
	return []*lib.RequestHandler{indexHandler}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from main controller")
}
