package main

import "fmt"

func main() {

	var x int
	fmt.Println("Qual o numero?")
	fmt.Scanln(&x)

	fmt.Println("O antecessor de ", x, " é ", x-1)
	fmt.Println("O sucessor de ", x, "é ", x+1)

}
