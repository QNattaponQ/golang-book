package main

import "fmt"

func main()  {
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
}