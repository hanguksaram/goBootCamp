package main

import "net/http"

func main() {

	startup := &Startup{}
	startup.Initialize()
	http.ListenAndServe(":8080", nil)
}

//client should implement logic when middleware stops to execute handlers chain

// func applyMiddleware(handlers map[handleRoute]http.HandlerFunc, middlewares ...Middleware) map[handleRoute]http.HandlerFunc {
// 	for key, hn := range handlers {
// 		for index := len(middlewares) - 1; index >= 0; index-- {
// 			hn = middlewares[index](hn)
// 		}
// 		handlers[key] = hn
// 	}

// 	return handlers
// }
