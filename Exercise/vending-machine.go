package main

import "fmt"

type NewVendingMachine struct {
	coins string
	items string
}

func (c NewVendingMachine) InsertCoin(x string) string {
	y := map[string]string{
		"T":  "10",
		"F":  "5",
		"TW": "2",
		"O":  "1",
	}
	return y[x]
}

func (i NewVendingMachine) SelectSD() string {
	return "Soft drink"
}

func main() {
	vm := NewVendingMachine{coins,items}
	
	//vm.InsertCoin("F")
	//vm.InsertCoin("TW")
	//vm.InsertCoin("O")
	//fmt.Println("Inserted Money:", vm.GetInsertedMoney()) // 18
	fmt.Println(vm.InsertCoin("F"))
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