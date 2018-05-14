package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type AutoGenerated struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp     float64 `json:"temp"`
		Pressure int     `json:"pressure"`
		Humidity int     `json:"humidity"`
		TempMin  int     `json:"temp_min"`
		TempMax  int     `json:"temp_max"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int     `json:"type"`
		ID      int     `json:"id"`
		Message float64 `json:"message"`
		Country string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	} `json:"sys"`
	ID   int    `json:"id"`
	Name string `json:"name"`
	Cod  int    `json:"cod"`
}

func main() {
	var port string
	flag.StringVar(&port, "port", ":3000", "default port:3000")
	flag.Parse()

	http.ListenAndServe(port, NewRouter())
}

func GETALLCITY(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	city := vars["city"]

	allCity := [5]string{"bangkok", "hobart", "nairobi", "newyork", "kupang"}
	for i := 0; i < 5; i++ {
		city = allCity[i]
		weatherReturn := GetWeather(city)
		fmt.Fprintf(w, "%s\n", weatherReturn.Name)
		fmt.Fprintf(w, "%vc ", weatherReturn.Main.Temp)
		fmt.Fprintf(w, "%s ", weatherReturn.Weather[0].Main)
		fmt.Fprintf(w, "%s\n\n", weatherReturn.Weather[0].Description)

	}

}

func GETBYCITY(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	city := vars["city"]

	weatherReturn := GetWeather(city)
	fmt.Fprintf(w, "%s\n", weatherReturn.Name)
	fmt.Fprintf(w, "%vc ", weatherReturn.Main.Temp)
	fmt.Fprintf(w, "%s ", weatherReturn.Weather[0].Main)
	fmt.Fprintf(w, "%s\n\n", weatherReturn.Weather[0].Description)

}

func GetWeather(city string) (genWeather *AutoGenerated) {
	stubURL := "http://localhost:8882/api/v1/weather/"
	res, _ := http.Get(stubURL + city)
	genWeather = new(AutoGenerated)
	json.NewDecoder(res.Body).Decode(genWeather)
	return
}

/*
func PrintOutPut(weatherReturn *AutoGenerated) {
	name := weatherReturn.Name
	temp := weatherReturn.Main.Temp
	weathermain := weatherReturn.Weather[0].Main
	weatherDesc := weatherReturn.Weather[0].Description

	return
}
*/
func NewRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/weather/all", GETALLCITY).Methods("GET")
	r.HandleFunc("/weather/{city}", GETBYCITY).Methods("GET")

	return r
}