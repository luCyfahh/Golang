package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	m := make(map[string]string)

	var name string
	var address string

	in := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter your name: ")
	in.Scan()
	name = in.Text()

	fmt.Print("Enter your address: ")
	in.Scan()
	address = in.Text()

	m["name"] = name
	m["address"] = address

	jsonObj, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jsonObj))
}
