package main

import "fmt"

func main() {
	var x map[string] int
	x = make(map[string]int)
	x["key"] = 10
	fmt.Println(x)
	fmt.Println(x["key"])

	y:= map[string]int {
		"two": 2,
		"one": 1,
		"Three": 3,

	}
	fmt.Println(y)
	fmt.Println(y["one"])
}