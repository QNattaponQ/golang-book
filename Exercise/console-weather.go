package main

import "fmt"

func main() {
	fmt.Println(weatherCelsius(25, "Bankok few cloud"))
	fmt.Println(weatherCelsius(34, "Tak sunny"))
	fmt.Println(weatherCelsius(17, "Phuket rainy"))
	fmt.Println(weatherCelsius(9, "Chiang-mai cold"))
}

func weatherCelsius(temp int, show string) string {
	if temp == 25 {
		fmt.Println(" _   _  ")
		fmt.Println(" _| |_  ")
		fmt.Println("|_   _|  c")

	} else if temp == 34 {
		fmt.Println(" _     ")
		fmt.Println(" _| |_|  ")
		fmt.Println(" _|   |  c")
	}
	return show
}
