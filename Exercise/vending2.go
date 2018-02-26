package main

import (
	"fmt"
)

type VendingMachine struct {
	insertedMoney int
	coins map[string]int
	items map[string]int
}

func (m VendingMachine) InsertedMoney() int {
	return m.insertedMoney
}

func (m *VendingMachine) InsertCoin(coin string) {
	m.insertedMoney+=m.coins[coin]
}

func (m *VendingMachine) SelectSD() string {
	price := m.items["SD"]
	change := m.insertedMoney-price
	m.insertedMoney=0
	return "Soft Drink" + m.change(change)
}

func (m *VendingMachine) SelectCC() string {
	price := m.items["CC"]
	change := m.insertedMoney - price
	m.insertedMoney=0
	return "Can Coffee" + m.change(change)
}

func (m *VendingMachine) SelectDK() string {
	price := m.items["DK"]
	change := m.insertedMoney - price
	m.insertedMoney=0
	return "Drinking water" + m.change(change)
	
}

func (m VendingMachine) change(c int) string {
	var str string
	if c==8 {
		str += ", F, TW, O"
	}
	return str
}


func main() {
	var cfcoins = map[string]int{
		"T":  10,
		"F":  5,
		"TW": 2,
		"O":  1,
	}

	var cfitems = map[string]int {
		"SD": 18,
		"CC": 12,
		"DW": 7,
	}
	vm := VendingMachine{coins: cfcoins,items :cfitems}
	fmt.Println("Currently inserted money:", vm.InsertedMoney())
	vm.InsertCoin("T")
	vm.InsertCoin("F")
	vm.InsertCoin("TW")
	vm.InsertCoin("O")
	fmt.Println("Currently inserted money:", vm.InsertedMoney())
	can := vm.SelectSD()
	fmt.Println(can)	
	
	vm.InsertCoin("T")
	vm.InsertCoin("T")
	fmt.Println("Currently inserted money:", vm.InsertedMoney())
	can = vm.SelectCC()
	fmt.Println(can)
	
	vm.InsertCoin("T")
	fmt.Println("Currently inserted money:", vm.InsertedMoney())
	can = vm.SelectDK()
	fmt.Println(can)
}