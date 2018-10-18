package lib

import "net/http"

type Middleware func(http.HandlerFunc) http.HandlerFunc
type MiddlewareChunk func(w http.ResponseWriter, r *http.Request) bool
type RequestHandler struct {
	Route         string
	Handler       http.HandlerFunc
	IsMiddlewared bool
}
type Controller interface {
	InitController() []*RequestHandler
}

func CreateNewMiddleware(middleWare func(w http.ResponseWriter, r *http.Request) bool) Middleware {

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
func ApplyMiddleware(handlers []*RequestHandler, middlewares ...Middleware) []*RequestHandler {
	for key, hn := range handlers {
		for index := len(middlewares) - 1; index >= 0; index-- {
			if hn.IsMiddlewared {
				hn.Handler = middlewares[index](hn.Handler)
			}

		}
		handlers[key] = hn
	}

	return handlers
}
