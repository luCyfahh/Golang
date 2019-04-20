package main

import (
	"fmt"
)

func swap(n int, input ...int) {
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			if input[i-1] > input[i] {
				input[i], input[i-1] = input[i-1], input[i]
				swapped = true
			}
		}
	}
	fmt.Println(input)
}

func bubbleSort(input ...int) {
	swap(len(input), input...)
}

func main() {
	n := 0
	x := make([]int, 0)
	fmt.Printf("Please input 10 random numbers\n")
	for i := 0; i < 10; i++ {
		fmt.Scan(&n)
		x = append(x, n)
	}
	bubbleSort(x...)
}
