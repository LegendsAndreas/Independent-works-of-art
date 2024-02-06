package main

import (
	"fmt"
	"log"
)

// All the needed methods for one type, that can work with the animal struct.
type Animal interface {
	Eat()
	Move()
	Speak()
}

// Structure for an animal
type animal struct {
	food       string
	locomotion string
	noise      string
}

func main() {
	// Variables for the user to assign to.
	var command, stringOne, stringTwo string

	// Creates the animals: Cow, bird and snake
	cow := animal{"grass", "walk", "moo"}
	bird := animal{"worms", "fly", "peep"}
	snake := animal{"mice", "slither", "hsss"}

	// Creates interfaces for the cow, bird and snake type.
	var interCow Animal = cow
	var interBird Animal = bird
	var interSnake Animal = snake

	// Map for animal name entered by the user for stringOne, in case the command is == newanimal, along with what type of animal it is (cow, bird or snake).
	var animalMap map[string]Animal
	animalMap = make(map[string]Animal)

	// Continues until the user enters x.
	for {
		fmt.Printf(">")
		_, err := fmt.Scan(&command, &stringOne, &stringTwo)

		if command == "x" || stringOne == "x" || stringTwo == "x" {
			break
		}

		// In case of an error, the program ends with the log of the error.
		if err != nil {
			log.Fatal(err)
		}

		// If the wants to create a new animal.
		if command == "newanimal" {
			switch stringTwo {
			case "cow":
				animalMap[stringOne] = interCow
				println("Created it!")
			case "bird":
				animalMap[stringOne] = interBird
				println("Created it!")
			case "snake":
				animalMap[stringOne] = interSnake
				println("Created it!")
			default:
				print("Incompatible type of animal")
			}

		} else if command == "query" { // If the user wants to see an association of an entered animal.
			// Sends the Animal interface of the associated, user entered name and which action to use.
			actionFunc(animalMap[stringOne], stringTwo)
		} else { // If the command is not equal to "newanimal" or "query", it prints wrong command
			print("Wrong command.")
		}
	}
}

// Methods for food, locomotion and noise.
func (ani animal) Eat() {
	fmt.Println(ani.food)
}

func (ani animal) Move() {
	fmt.Println(ani.locomotion)
}

func (ani animal) Speak() {
	fmt.Println(ani.noise)
}

// Function for which action to use.
func actionFunc(a Animal, act string) {
	if act == "eat" {
		a.Eat()
	} else if act == "move" {
		a.Move()
	} else if act == "speak" {
		a.Speak()
	} else {
		print("Action is not compatible")
	}
}
