package main

import (
	"fmt"
	"log"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func main() {
	// Creates the animals: Cow, bird and snake
	cow := Animal{"grass", "walk", "moo"}

	var bird Animal
	bird.food = "worms"
	bird.locomotion = "fly"
	bird.noise = "peep"

	var snake Animal
	snake.food = "mice"
	snake.locomotion = "slither"
	snake.noise = "hsss"

	var creature string
	var action string

	for {
		// Asks for which animal and action to select and handles error. You only need to create 1 err variable, if you want to assign to multiple variables.
		fmt.Print(">")
		_, err := fmt.Scan(&creature, &action)

		if err != nil {
			log.Fatal(err)
		}

		// If creature is equal to cow.
		if creature == "cow" {
			if action == "food" {
				cow.Eat()
			} else if action == "locomotion" {
				cow.Move()
			} else if action == "noise" {
				cow.Speak()
			}
		}

		// If creature is equal to bird
		if creature == "bird" {
			if action == "food" {
				bird.Eat()
			} else if action == "locomotion" {
				bird.Move()
			} else if action == "noise" {
				bird.Speak()
			}
		}

		// If creature is equal to snake
		if creature == "snake" {
			if action == "food" {
				snake.Eat()
			} else if action == "locomotion" {
				snake.Move()
			} else if action == "noise" {
				snake.Speak()
			}
		}
	}

}

// Methods for food, locomotion and noise.
func (ani Animal) Eat() {
	fmt.Println(ani.food)
}

func (ani Animal) Move() {
	fmt.Println(ani.locomotion)
}

func (ani Animal) Speak() {
	fmt.Println(ani.noise)
}
