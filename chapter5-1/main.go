package main

import "fmt"

func main() {
	/*
		fmt.Println("1")
		fmt.Println("2")
		fmt.Println("3")
		fmt.Println("4")
		fmt.Println("5")
		fmt.Println("6")
		fmt.Println("7")
		fmt.Println("8")
		fmt.Println("9")
		fmt.Println("10")

		//Condition For
		i := 1
		for i <= 10 {
			fmt.Println(i)
			i = i + 1
		}
	
	
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
