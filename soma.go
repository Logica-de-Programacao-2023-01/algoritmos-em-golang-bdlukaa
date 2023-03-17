package main

import "fmt"

func main() {
	var x1, x2, x3 int

	fmt.Println("Number 1")
	fmt.Scan(&x1)

	fmt.Println("Number 2")
	fmt.Scan(&x2)

	fmt.Println("Number 3")
	fmt.Scan(&x3)

	fmt.Println("The sum of the three numbers is ", x1+x2+x3)

}
