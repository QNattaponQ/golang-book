package main

import "fmt"

func main()  {
	//Switch Case
	//var j int
	for j := 1; j <=100; j++ {
		switch {
		case j%15==0: fmt.Println(j, "FizzBuzz")
		case j%3==0: fmt.Println(j, "Fizz")
		case j%5==0: fmt.Println(j, "Buzz")
		default: fmt.Println(j)
			
		}
	}
}