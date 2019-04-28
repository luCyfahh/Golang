package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var inputStr string
	var intV int

	slcX := make([]int, 3)

me:
	for i := 0; strings.ToLower(inputStr) != "x"; i++ {
		fmt.Printf("Please enter some integers? ")
		fmt.Scan(&inputStr)

		if strings.ToLower(inputStr) != "x" {
			intV, _ = strconv.Atoi(inputStr)

			if i < 3 {
				slcX[i] = intV
			} else {
				slcX = append(slcX, intV)
			}

			slice := make([]int, len(slcX))
			copy(slice, slcX)
			sort.Ints(slice)
			fmt.Println(slice)
		} else {
			break me
		}
	}
}
