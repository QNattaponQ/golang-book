package main

import "fmt"

func main() {
	x:= map[string]int {
		"one": 1,
		"two": 2,
		"Three": 3,

	}
	fmt.Println(x)

	delete(x, "two")
	fmt.Printf("After delete: %v\n", x)
}