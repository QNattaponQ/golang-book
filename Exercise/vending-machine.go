package main

import "fmt"

type NewVendingMachine struct {
	coins string
	items string
}

func (c NewVendingMachine) InsertCoin() string {
	return "10"
}

func (i NewVendingMachine) SelectSD() string {
	var output string

	if i == "SD" {
		output = "Soft drink"
	}
	
	return output
}

func main() {
	vm := NewVendingMachine{"T", "SD"}
	fmt.Println(vm.InsertCoin())
	//vm.InsertCoin("F")
	//vm.InsertCoin("TW")
	//vm.InsertCoin("O")
	//fmt.Println("Inserted Money:", vm.GetInsertedMoney()) // 18
	can := vm.SelectSD()
	fmt.Println(can) // SD
}

/*
func GetInsertedMoney() int {
	coins := map[string]int{
		"T":  10,
		"F":  5,
		"TW": 2,
		"O":  1,
	}
	
}
*/