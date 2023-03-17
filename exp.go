package main

import "fmt"

func main() {
	var x1 int

	fmt.Println("What is the number, please?")
	fmt.Scanln(&x1)

	fmt.Println("Its double is ", x1*2)
	fmt.Println("Its triple is ", x1*3)
	fmt.Println("Its quadruple is ", x1*4)
}
