package main

import "fmt"

func main() {
	fmt.Println("Holla")

	arr := [...]int{12, 40, 23, 44, 30, 18, 27}
	//arr := [...]int{1, 2, 3, 4}
	var sum float64
	for x := 0; x < len(arr); x++ {
		sum = sum + float64(arr[x])
	}

	fmt.Println("la moyen est =>", sum/float64(len(arr)))
}
