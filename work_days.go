package main

import "fmt"

func main() {

	var workedDays int
	var dayPrice float64

	fmt.Println("How many days have you worked?")
	fmt.Scan(&workedDays)

	fmt.Println("How much does your day cost?")
	fmt.Scan(&dayPrice)

	var wage = float64(workedDays) * dayPrice
	fmt.Println("You end wage is ", wage)

}
