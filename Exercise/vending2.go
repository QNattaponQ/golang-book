package main

import (
	"fmt"
)

type VendingMachine struct {
	insertedMoney int
}

func (m VendingMachine) InsertedMoney() int {
	return m.insertedMoney
}

func (m *VendingMachine) InsertCoin(coin string) {
	m.insertedMoney=10
}


func main() {
	vm := VendingMachine{}
	fmt.Println("Currently inserted money:", vm.InsertedMoney())
	vm.InsertCoin("T")
	fmt.Println("Currently inserted money:", vm.InsertedMoney())	
}