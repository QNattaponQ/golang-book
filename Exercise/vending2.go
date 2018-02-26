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

func (m *VendingMachine) SelectSD() string {
	m.insertedMoney=0
	return "Soft Drink"
}

func (m *VendingMachine) SelectCC() string {
	m.insertedMoney=0
	return "Can Coffee"
}

func (m *VendingMachine) SelectDK() string {
	m.insertedMoney=0
	return "Drinking water"
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
	
	vm.InsertCoin("T")
	vm.InsertCoin("TW")
	fmt.Println("Currently inserted money:", vm.InsertedMoney())
	can = vm.SelectCC()
	fmt.Println(can)
	
	vm.InsertCoin("T")
	vm.InsertCoin("T")
	fmt.Println("Currently inserted money:", vm.InsertedMoney())
	can = vm.SelectDK()
	fmt.Println(can)
}