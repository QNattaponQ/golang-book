package main

import (
	"fmt"
	"net/http"
)

type HomePageHandler struct{}

func (h *HomePageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")  //localhost:3000/?name=xyz
	if name == "" {
		name = "World"
	}
	fmt.Fprintln(w, "Hello,", name)
}

func main() {
	http.Handle("/", &HomePageHandler{})

	http.ListenAndServe(":3000", nil)
}
