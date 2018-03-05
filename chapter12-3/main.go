package main

import (
	"fmt"
	"strconv"
)

func main() {
	// addFunc := func(a int) (func(b int) int ) {
	// 	return func(b int) int {
	// 		return a + b
	// 	}
	// }
	// // addNewFunc := addFunc(2)
	// // fmt.Println(addNewFunc(3))
	// add2With := addFunc(2)
	// fmt.Println(add2With(3))
	// for i:=1; i <100; i++ {
	// 	fmt.Println(
	// 		func(a int) string {
	// 		if a%15==0 {
	// 			return "FizzBuzz"
	// 		}
	// 		if a%3 == 0 {
	// 			return "Fizz"
	// 		}
	// 		if a%5==0 {
	// 			return "Buzz"
	// 		}
	// 		return strconv.Itoa(a)
	// 	} (i))
	// }
	for j := 1; j <= 100; j++ {
		fmt.Println(j, findFizzBuzz(j))
	}
}

func findFizzBuzz(number int) string {

	// fizzBuzzFunc := func(n int) (string, bool) {
	// 	if n%15==0 {
	// 		return "FizzBuzz", true
	// 	}
	// 	return "", false
	// }

	// fizzFunc := func(n int) (string, bool) {
	// 	if n%3==0 {
	// 		return "Fizz", true
	// 	}
	// 	return "", false
	// }
	// buzzFunc := func(n int) (string, bool) {
	// 	if n%5==0 {
	// 		return "Buzz", true
	// 	}
	// 	return "", false
	// }

	fbTemplate := func(fbnumber int, str string) func(int) (string, bool) {
		return func(n int) (string, bool) {
			if n%fbnumber == 0 {
				return str, true
			}
			return "", false
		}
	}

	fbArray := [...]func(n int) (string, bool){
		fbTemplate(15, "FizzBuzz"),
		fbTemplate(5, "Buzz"),
		fbTemplate(3, "Fizz"),
	}

	for i := 0; i < len(fbArray); i++ {
		if str, ok := fbArray[i](number); ok {
			return str
		}
	}
	return strconv.Itoa(number)
}
