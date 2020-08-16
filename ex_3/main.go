package main

import "fmt"

type Address struct {
	name    string
	street  string
	city    string
	state   string
	Pincode int
}

func main() {
	var address Address
	address = Address{name: "Auberge", street: "Mutsanga", city: "Butembo", state: "DRC", Pincode: 23}
	fmt.Println(address)
	// for _, t := range address {
	// 	fmt.Println(t)
	// }
}
