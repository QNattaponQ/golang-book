package main

import (
	"fmt"
	"math/rand"
)

func main() {
	random := rand.ExpFloat64()

	for i := 1; i <= 10; i++ {
		fmt.Print("Enter a number: ")
		var input float64
		fmt.Scanf("%f\n", &input) //Prompt Input data

		if i > 5 {
			fmt.Println("Enough")
			{
				break
			}
		}

		if input == random {
			fmt.Println(input, random, "Match")
		} else if input < random {
			fmt.Println(input, random, "Input less than random")
		} else if input > random {
			fmt.Println(input, random, "Input more than random")
		}

	}

}
