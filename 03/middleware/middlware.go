package middleware

import (
	"fmt"
	"net/http"

	"github.com/GoesToEleven/udemyTraining/03/lib"
)

type MiddlewareFactory struct{}

func (mf *MiddlewareFactory) PrepareMiddleware(mw lib.MiddlewareChunk) lib.Middleware {
	return lib.CreateNewMiddleware(mw)
}

func AllowCors(w http.ResponseWriter, r *http.Request) bool {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	return true
}
func CustomMiddleWare1(w http.ResponseWriter, r *http.Request) bool {
	fmt.Fprintln(w, "CUstom middleware1")
	return true
}
func CustomMiddleWare2(w http.ResponseWriter, r *http.Request) bool {
	fmt.Fprintln(w, "CUstom middleware2")
	return false
}
