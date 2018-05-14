package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("recover")
		if p := recover(); p != nil {
			fmt.Printf("Panic found %v\n", p)
		}
	}()

	f()

}

func f() {
	defer func() {
		fmt.Println("Deferred by f")
	}()
	g()
}

func g() {
	defer func() {
		fmt.Println("Deferred by g")
	}()
	panic("AAA")
	h()

}

func h() {
	defer func() {
		fmt.Println("Deferred by h")
	}()

}
