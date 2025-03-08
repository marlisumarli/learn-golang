package main

import (
	"fmt"
	"log"
)

func main() {
	var input float32

	fmt.Print("Enter float number: ")
	_, err := fmt.Scan(&input)
	if err != nil {
		log.Fatal(err)
	}

	intNumber := int(input)

	fmt.Println("Number converted:", intNumber)
}
