package main

import "fmt"

func main() {

	var x1, x2 int

	fmt.Println("Number 1")
	fmt.Scanln(&x1)

	fmt.Println("Number 2")
	fmt.Scanln(&x2)

	if x1 == x2 {
		fmt.Println("Equal numbers")
	} else if x1 < x2 {
		fmt.Println(x1, " is smaller than ", x2)
	} else {
		fmt.Println(x2, " is smaller than ", x1)
	}

}
