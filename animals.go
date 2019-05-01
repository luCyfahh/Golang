package main

import (
	"fmt"
	"strings"
)

func main() {
	aStr := ""
	bStr := ""

	cow := Animal{
		food:       "grass",
		locomotion: "walk",
		noise:      "moo",
	}

	bird := Animal{
		food:       "worms",
		locomotion: "fly",
		noise:      "peep",
	}

	snake := Animal{
		food:       "mice",
		locomotion: "slither",
		noise:      "hsss",
	}

	for {
		fmt.Print(">")
		fmt.Scan(&aStr, &bStr)

		aStr = strings.ToLower(aStr)
		bStr = strings.ToLower(bStr)

		if aStr == "cow" {
			if bStr == "eat" {
				cow.Eat()
			}

			if bStr == "speak" {
				cow.Speak()
			}

			if bStr == "move" {
				cow.Move()
			}
		}

		if aStr == "bird" {
			if bStr == "eat" {
				bird.Eat()
			}

			if bStr == "speak" {
				bird.Speak()
			}

			if bStr == "move" {
				bird.Move()
			}
		}

		if aStr == "snake" {
			if bStr == "eat" {
				snake.Eat()
			}

			if bStr == "speak" {
				snake.Speak()
			}

			if bStr == "move" {
				snake.Move()
			}
		}
	}
}

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() {
	fmt.Println("Type of food: ", a.food)
}
func (a Animal) Speak() {
	fmt.Println("Type of sound: ", a.noise)
}
func (a Animal) Move() {
	fmt.Println("Type of locomotion: ", a.locomotion)
}
