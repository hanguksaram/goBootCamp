package middleware

import (
	"fmt"
	"net/http"
)

func CustomMiddleWare1(w http.ResponseWriter, r *http.Request) bool {
	fmt.Fprintln(w, "CUstom middleware1")
	return true
}
func CustomMiddleWare2(w http.ResponseWriter, r *http.Request) bool {
	fmt.Fprintln(w, "CUstom middleware2")
	return false
}
