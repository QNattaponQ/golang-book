package main

import "fmt"

func main() {
	/*
		//Condition IF
		for j := 1; j <= 100; j++ {
			if j%15 == 0 {
				fmt.Println(j, "FizzBuzz")
			} else if j%3 == 0 {
				fmt.Println(j, "Fizz")
			} else if j%5 == 0 {
				fmt.Println(j, "Buzz")
			} else {
				fmt.Println(j)
			}
		}
	*/
	for j := 1; j <= 100; j++ {
		fmt.Println(j, findFizzBuzz(j))
	}
}

func findFizzBuzz(j int) (result string) {
	if j%5 == 0 {
		result = "Buzz"
	} else if j%3 == 0 {
		result = "Fizz"
	} else if j%15 == 0 { 
		result = "FizzBuzz"
	}
	return
}
