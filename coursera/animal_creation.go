package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}

func (c Cow) Eat()   { fmt.Println("grass") }
func (c Cow) Move()  { fmt.Println("walk") }
func (c Cow) Speak() { fmt.Println("moo") }

type Bird struct{}

func (b Bird) Eat()   { fmt.Println("worms") }
func (b Bird) Move()  { fmt.Println("fly") }
func (b Bird) Speak() { fmt.Println("peep") }

type Snake struct{}

func (s Snake) Eat()   { fmt.Println("mice") }
func (s Snake) Move()  { fmt.Println("slither") }
func (s Snake) Speak() { fmt.Println("hsss") }

func main() {
	animals := make(map[string]Animal)

	fmt.Println("Commands:")
	fmt.Println("1. newanimal <name> <type>   - Create a new animal (type: cow, bird, snake)")
	fmt.Println("2. query <name> <action>     - Query an animal (action: eat, move, speak)")
	fmt.Println("Type 'exit' to quit.")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")

		scanner.Scan()
		input := strings.ToLower(strings.TrimSpace(scanner.Text()))

		if input == "exit" {
			fmt.Println("Exiting program.")
			break
		}

		parts := strings.Split(input, " ")
		if len(parts) != 3 {
			fmt.Println("Invalid input format. Use: newanimal <name> <type> OR query <name> <action>")
			continue
		}

		command, name, param := parts[0], parts[1], parts[2]

		switch command {
		case "newanimal":
			var newAnimal Animal

			switch param {
			case "cow":
				newAnimal = Cow{}
			case "bird":
				newAnimal = Bird{}
			case "snake":
				newAnimal = Snake{}
			default:
				fmt.Println("Invalid animal type. Choose from: cow, bird, snake.")
				continue
			}

			animals[name] = newAnimal
			fmt.Println("Created it!")

		case "query":
			animal, exists := animals[name]
			if !exists {
				fmt.Println("Animal not found.")
				continue
			}

			switch param {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Println("Invalid query. Use: eat, move, speak.")
			}

		default:
			fmt.Println("Unknown command. Use 'newanimal' or 'query'.")
		}
	}
}
