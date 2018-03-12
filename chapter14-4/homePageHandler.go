package home

import (
	"fmt"
	"net/http"
)

func HomePageHandler (res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello world!")
}