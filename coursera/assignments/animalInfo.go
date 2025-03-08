package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Define the Animal type as a struct
type Animal struct {
	food       string
	locomotion string
	noise      string
}

// Define the Eat method for the Animal type
func (a Animal) Eat() {
	fmt.Println(a.food)
}

// Define the Move method for the Animal type
func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

// Define the Speak method for the Animal type
func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	// Create a map to store the predefined animals
	animals := map[string]Animal{
		"cow":   {food: "grass", locomotion: "walk", noise: "moo"},
		"bird":  {food: "worms", locomotion: "fly", noise: "peep"},
		"snake": {food: "mice", locomotion: "slither", noise: "hsss"},
	}

	// Create a scanner to read user input
	scanner := bufio.NewScanner(os.Stdin)

	// Start an infinite loop to handle user requests
	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()

		// Split the input into two parts: animal and action
		parts := strings.Fields(input)
		if len(parts) != 2 {
			fmt.Println("Invalid input. Please enter an animal and an action (e.g., 'cow eat').")
			continue
		}

		animalName := parts[0]
		action := parts[1]

		// Check if the animal exists in the map
		animal, exists := animals[animalName]
		if !exists {
			fmt.Println("Animal not found. Available animals: cow, bird, snake.")
			continue
		}

		// Call the appropriate method based on the action
		switch action {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Println("Invalid action. Available actions: eat, move, speak.")
		}
	}
}
