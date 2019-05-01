package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (animal *Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal *Animal) Move() {
	fmt.Println(animal.locomotion)
}

func (animal *Animal) Speak() {
	fmt.Println(animal.noise)
}

func main() {

	cow := Animal{food: "grass", locomotion: "walk", noise: "moo"}
	bird := Animal{food: "worms", locomotion: "fly", noise: "peep"}
	snake := Animal{food: "mice", locomotion: "slither", noise: "hss"}

	var input string

	for !strings.Contains(input, "exit") {
		fmt.Printf("Enter an animal and action (two words. 'Exit' for exit): ")

		//read input and convert it to all low case
		in := bufio.NewReader(os.Stdin)
		input, _ = in.ReadString('\n')
		input = strings.ToLower(input)

		//split the word by space
		words := strings.Split(input, " ")

		//default animal is snake in case input doesn't match
		selectedAnimal := snake

		if len(words) == 2 {

			//assign selected animal based on input
			switch words[0] {
			case "cow":
				selectedAnimal = cow
			case "bird":
				selectedAnimal = bird
			default:
				selectedAnimal = snake
			}

			//call method based on input. Default action is speak if input is wrong
			if strings.Contains(words[1], "eat") {
				selectedAnimal.Eat()
			} else if strings.Contains(words[1], "move") {
				selectedAnimal.Move()
			} else {
				selectedAnimal.Speak()
			}
		} else {
			fmt.Println("WRONG INPUT!! Please enter TWO WORDS!!")
		}

	}
}
