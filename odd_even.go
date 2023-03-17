package main

import "fmt"

func main() {

	var x1 int

	fmt.Println("Number")
	fmt.Scanln(&x1)

	if x1%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

}
