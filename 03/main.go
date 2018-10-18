package main

import "net/http"

func main() {

	startup := &Startup{}
	startup.Initialize()
	http.ListenAndServe(":8080", nil)
}
