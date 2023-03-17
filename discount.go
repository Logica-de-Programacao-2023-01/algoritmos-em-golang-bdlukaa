package main

import "fmt"

func main() {

	var current float64

	fmt.Println("Valor do produto?")
	fmt.Scan(&current)

	var percent = current * 0.10

	fmt.Println("O valor com desconto é", current-percent)

}
