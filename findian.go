package main

import (
	"fmt"
	"strings"
)

func main() {
	var strN string

	fmt.Printf("Please enter a string? ")
	fmt.Scan(&strN)

	if strings.HasPrefix(strN, "i") && strings.HasSuffix(strN, "n") && strings.Contains(strN, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
