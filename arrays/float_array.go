package main

import "fmt"

func main() {
	var numeros = [6]float32{1, 2, 3, 4, 5, 6}

	var media float32 = 0

	for _, numero := range numeros {
		media += numero
	}

	fmt.Println("A média dos números é de", media/float32(len(numeros)))
}
