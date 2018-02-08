package main

import (
	"fmt"
)

func main() {
	//Long Declaration
	var x string = "Hello world"
	fmt.Println(x)

	var y string
	y = "Hello world"
	fmt.Println(y)

	//Short Declaration
	//Type Inference
	z := "Hello world"
	fmt.Println(z)
	fmt.Printf("Type: %T\n", z)
}
