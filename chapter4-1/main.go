package main

import (
	"fmt"
)

func main() {
	//Long Declaration
	var x string = "Hello world"
	fmt.Println(x)
	fmt.Printf("Type: %T\n", x)

	var y string
	y = "Hello world"
	fmt.Println(y)
	fmt.Printf("Type: %T\n", y)

	//Short Declaration
	//Type Inference
	z := "Hello world"
	fmt.Println(z)
	fmt.Printf("Type: %T\n", z)

	q := 1342.01
	fmt.Println(q)
	fmt.Printf("Type: %T\n", q)

	//Constant assign static value
	const xxx string = "Hello, world"
	//b = "Other string"

	//Set variable
	var (
		a = 5
		b = 10
		c = 15
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	v1, v2 := "first", "second"
	v1, v2 = v2, v1 //Swap variable
	fmt.Println(v1)
	fmt.Println(v2)

}
