package main

import "fmt"

func main() {
	fmt.Println("Quantos anos você têm?")
	var age float64
	fmt.Scan(&age)

	var days float64 = age * 365.3
	fmt.Println("Você tem ", days, " dias")

}
