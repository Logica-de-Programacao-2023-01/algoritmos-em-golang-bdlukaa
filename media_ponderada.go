package main

import "fmt"

func main() {

	var x1, x2, x3 float64

	fmt.Println("Diga os numeros:")
	fmt.Scanln(&x1)
	fmt.Scanln(&x2)
	fmt.Scanln(&x3)

	var media = ((x1 * 2) + (x2 * 3) + (x3 * 5)) / 3

	fmt.Println("A média ponderada dos números é ", media)

}
