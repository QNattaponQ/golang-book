package main

import "fmt"

func main() {
	numbers := [5] int {1,2,3,4,5}
	for i:=0; i< len(numbers); i++ {
		fmt.Println(i,numbers[i])
	}

	fmt.Println("White Range")
	for i,numbers := range numbers {
		fmt.Println(i,numbers)
	}
}