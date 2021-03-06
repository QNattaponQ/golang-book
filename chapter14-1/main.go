package main

import (
	"fmt"
	"net/http"
)

type HomePageHandler struct {}

func (h *HomePageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world")
}

func main() {
	http.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Hello world")
		})
	barHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello bar")
	}

	http.HandleFunc("/bar", barHandler)

	http.Handle("/home", &HomePageHandler{})

	http.ListenAndServe(":3000",nil)
}
