package main

import "fmt"

func main() {

	var x1 int

	fmt.Println("Number")
	fmt.Scanln(&x1)

	if x1%3 == 0 && x1%5 == 0 {
		fmt.Println("Multiple of 3 and 5")
	} else {
		fmt.Println("Not a multiple of 3 and 5")
	}

}
