package main

import "fmt"

func main() {

	var weight, height float64

	fmt.Println("What is the weight, please?")
	fmt.Scanln(&weight)

	fmt.Println("What is the height, please?")
	fmt.Scanln(&height)

	var imc = weight * (height * height)
	fmt.Println("Your IMC is ", imc)

}
