package main

import (
	"fmt"
)

type VendingMachine struct {
	coins string
}

func (m VendingMachine) InsertedMoney() int {
	return 0
}


func main() {
	vm := VendingMachine{}
	fmt.Println("Currently inserted money:", vm.InsertedMoney())	
}