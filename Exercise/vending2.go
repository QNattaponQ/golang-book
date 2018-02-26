package main

import (
	"fmt"
)

type VendingMachine struct {
	insertedMoney int
	coins map[string]int
}

func (m VendingMachine) InsertedMoney() int {
	return m.insertedMoney
}

func (m *VendingMachine) InsertCoin(coin string) {
	m.insertedMoney+=m.coins[coin]
}

func (m VendingMachine) SelectSD() string {
	return "Soft Drink"
}


func main() {
	var coins = map[string]int{
		"T":  10,
		"F":  5,
		"TW": 2,
		"O":  1,
	}
	vm := VendingMachine{coins: coins}
	fmt.Println("Currently inserted money:", vm.InsertedMoney())
	vm.InsertCoin("T")
	vm.InsertCoin("F")
	vm.InsertCoin("TW")
	fmt.Println("Currently inserted money:", vm.InsertedMoney())
	can := vm.SelectSD()
	fmt.Println(can)	
}