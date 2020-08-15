package main

import "fmt"

func main() {
	n := 12
	var x *int = &n
	*x = 13
	fmt.Println("Nous sommes la", n)
	fmt.Println("Nous sommes la", *x)
}
