package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func init() {

}

func main() {
	authCntrl := AuthController{}
	mux := httprouter.New()
	authCntrl.initAuthController(mux)
	http.ListenAndServe(":5000", mux)
}
