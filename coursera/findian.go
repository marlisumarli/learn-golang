package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Input any string: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Data: %s", input)

	input = strings.TrimSpace(input)

	lowerInput := strings.ToLower(input)

	if strings.HasPrefix(lowerInput, "i") && strings.HasSuffix(lowerInput, "n") && strings.Contains(lowerInput, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
