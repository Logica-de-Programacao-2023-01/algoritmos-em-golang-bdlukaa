package main

import "fmt"

func main() {

	var matriz [3][4]int

	for linha, _ := range matriz {

		for coluna, _ := range matriz[linha] {

			matriz[linha][coluna] = linha + coluna

		}

	}

	fmt.Println(matriz)

}
