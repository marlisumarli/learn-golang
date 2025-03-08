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

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	animals := map[string]Animal{
		"cow":   {"grass", "walk", "moo"},
		"bird":  {"worms", "fly", "peep"},
		"snake": {"mice", "slither", "hsss"},
	}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter a request in the format: <animal> <action> (e.g., 'cow eat')")
	fmt.Println("Available animals: cow, bird, snake")
	fmt.Println("Available actions: eat, move, speak")
	fmt.Println("Type 'exit' to quit.")

	for {
		fmt.Print("> ")

		scanner.Scan()
		input := strings.ToLower(strings.TrimSpace(scanner.Text()))

		if input == "exit" {
			fmt.Println("Exiting program.")
			break
		}

		parts := strings.Split(input, " ")
		if len(parts) != 2 {
			fmt.Println("Invalid input. Please enter in the format: <animal> <action>")
			continue
		}

		animalName, action := parts[0], parts[1]

		animal, exists := animals[animalName]
		if !exists {
			fmt.Println("Unknown animal. Available options: cow, bird, snake.")
			continue
		}

		switch action {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Println("Unknown action. Available options: eat, move, speak.")
		}
	}
}
