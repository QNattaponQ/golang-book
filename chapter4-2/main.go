package main

import (
	"fmt"
)

func main() {
	/*
		fmt.Print("Enter a number: ")
		var input float64
		fmt.Scanf("%f", &input) //Prompt Input data
		output := input * 2
		fmt.Println(output)
	*/

	//Problem solve convert Fahrenheit ot Celsius
	fmt.Print("Enter Fahrenheit: ")
	var fahrenheit float32
	var celsius float32
	fmt.Scanf("%f", &fahrenheit)
	celsius = ((fahrenheit - 32) * 5 / 9)
	fmt.Printf("Celsius : %.2f", celsius)
}
