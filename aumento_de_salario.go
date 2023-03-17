package main

import "fmt"

func main() {

	var current float64

	fmt.Println("Qual seu salario?")
	fmt.Scan(&current)

	var percent = current * 0.15

	fmt.Println("Seu novo salário é de", current+percent)

}
