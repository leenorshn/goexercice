package main

import "fmt"

type Address struct {
	name    string
	street  string
	city    string
	state   string
	Pincode int
}

type Car struct {
	Name, Model, Color string
	WeightInKg         float64
}

func main() {
	// var address Address
	// address = Address{name: "Auberge", street: "Mutsanga", city: "Butembo", state: "DRC", Pincode: 23}
	// c := Car{"Rav4", "Voiture", "Metalique", 5788.9}
	// fmt.Println(address)
	// fmt.Println(c.Color)
	// c.Model = "Toyota"
	// fmt.Println(c.Model)

	// for _, t := range address {
	// 	fmt.Println(t)
	// }

	c := &Car{"Rav4", "Voiture", "Metalique", 5788}
	fmt.Println(*c)
	c2 := &c
	(*c2).Model = "Hammer"
	fmt.Println((*c).Model)
}
