package main

import (
	"fmt"
	"math/rand"
)

func main() {
	/* Version 1
	random := rand..ExpFloat64()
	*/
	random := rand.Intn(100)

	for i := 1; i <= 10; i++ {
		fmt.Print("Enter a number: ")
		/* Version 1
		var input float64			
		fmt.Scanf("%f\n", &input)
		*/
		var input int
		fmt.Scanf("%d\n", &input) //Prompt Input data

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
