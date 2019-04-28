package main

import "fmt"

func main() {
	var floatN float64

	fmt.Printf("Please enter a float number? ")
	fmt.Scan(&floatN)
	fmt.Print("The int number is: ", int(floatN))
	fmt.Println()
}
