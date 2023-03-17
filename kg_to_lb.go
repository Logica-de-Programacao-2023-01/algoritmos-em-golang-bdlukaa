package main

import "fmt"

func main() {

	var current float64

	fmt.Println("Peso em KG?")
	fmt.Scan(&current)

	var lbs = current * 2.20462
	fmt.Println("O seu peso em libras é", lbs)

}
