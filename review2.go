package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	slice := make([]int, 3)
	reader := bufio.NewReader(os.Stdin)
	index := 0
	for {
		fmt.Println("Enter a digit:")
		userInput, err := reader.ReadString('\n')
		if err != nil {
			os.Exit(0)
		}
		if strings.ToLower(userInput) == "x\n" {
			break
		}
		userInput = strings.TrimRight(userInput, "\n")
		num, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Cannot convert input into <int>")
			continue
		}
		if index < 3 {
			slice[index] = num
			index++
		} else {
			slice = append(slice, num)
		}
		sliceCopy := make([]int, len(slice))
		copy(sliceCopy, slice)
		sort.Ints(sliceCopy)
		fmt.Println(sliceCopy)
	}
}
