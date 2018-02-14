package main

import (
	"fmt"
)

func main()  {
	fmt.Println(f1())
	fmt.Println(plus(5,5))
}

func f1() int {
	return f2()
}

func f2() int {
	return 1
}

func plus(a int, b int) int  {
	return a + b
}