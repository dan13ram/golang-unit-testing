package main

import (
	"fmt"
	"net/http"
)

type helloHandler struct{}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	if path == "" {
		path = "World"
	}
	fmt.Fprintf(w, "Hello, %s!", path)
}

func main() {
	http.ListenAndServe(":8080", &helloHandler{})
}
