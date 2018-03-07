package main

import (
	"fmt"
)

func main() {
	ch:= make(chan float64,10)
	go func() {ch <- 1} ()
	ch <-5
	ch <-9
	fmt.Println("cap", cap(ch)) //จำนวน Buffer ที่กำหนดไว้
	fmt.Println("len", len(ch)) //จำนวน ch ที่ถูก set ค่าไว้


}
