package main

import (
	"fmt"
	"github.com/QNattaponQ/vendingMachine"
)

var coins = map[string]int{
	"T":  10,
	"F":  5,
	"TW": 2,
	"O":  1,
}

var items = map[string]int {
	"SD": 18,
	"CC": 12,
	"DW": 7,
}

func main() {
	
	vm := vendingMachine.NewVendingMachine(coins,items)
	fmt.Println("Currently inserted money:", vm.InsertedMoney())
	vm.InsertCoin("T")
	vm.InsertCoin("F")
	vm.InsertCoin("TW")
	vm.InsertCoin("O")
	fmt.Println("Currently inserted money:", vm.InsertedMoney())
	can := vm.SelectSD()
	fmt.Println(can)	
	
	fmt.Println()
	vm.InsertCoin("T")
	vm.InsertCoin("T")
	fmt.Println("Currently inserted money:", vm.InsertedMoney())
	can = vm.SelectCC()
	fmt.Println(can)
	
	fmt.Println()
	vm.InsertCoin("T")
	fmt.Println("Currently inserted money:", vm.InsertedMoney())
	can = vm.SelectDK()
	fmt.Println(can)

	fmt.Println()
	vm.InsertCoin("T")
	vm.InsertCoin("F")
	fmt.Println("Currently inserted money:", vm.InsertedMoney())
	coin := vm.CoinReturn()
	fmt.Println(coin)
}