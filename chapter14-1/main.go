package main

import (
	"fmt"
	"net/http"
)

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

	http.ListenAndServe(":3000",nil)
}
