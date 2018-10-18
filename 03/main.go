package main

import (
	"fmt"
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc
type handleRoute string

var handlersStorare = make(map[handleRoute]http.HandlerFunc, 2)

func main() {

	handlersStorare["/"] = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world %v", 1)
	}
	handlersStorare["/test"] = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "test %v", 1)
	}

	middleWare1 := createNewMiddleware(customMiddleWare1)
	middleWare2 := createNewMiddleware(customMiddleWare2)

	applyMiddleware(handlersStorare, middleWare1, middleWare2)
	for key, hn := range handlersStorare {
		http.HandleFunc(string(key), hn)
	}
	http.ListenAndServe(":8080", nil)
}
func customMiddleWare1(w http.ResponseWriter, r *http.Request) bool {
	fmt.Fprintln(w, "CUstom middleware1")
	return false
}
func customMiddleWare2(w http.ResponseWriter, r *http.Request) bool {
	fmt.Fprintln(w, "CUstom middleware2")
	return false
}
func applyMiddleware(handlers map[handleRoute]http.HandlerFunc, middlewares ...Middleware) map[handleRoute]http.HandlerFunc {
	for key, hn := range handlers {
		for index := len(middlewares) - 1; index >= 0; index-- {
			hn = middlewares[index](hn)
		}
		handlers[key] = hn
	}

	return handlers
}
func createNewMiddleware(middleWare func(w http.ResponseWriter, r *http.Request) bool) Middleware {

	middleware := func(next http.HandlerFunc) http.HandlerFunc {

		handler := func(w http.ResponseWriter, r *http.Request) {
			if middleWare(w, r) {
				next(w, r)
			} else {
				return
			}

		}

		return handler
	}

	return middleware
}
