package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	numbers := make([]int, 0, 3)
	var input string

	for {
		fmt.Print("Enter number (Type 'X' to exit): ")
		_, err := fmt.Scan(&input)
		if err != nil {
			log.Fatal(err)
		}

		if strings.ToLower(input) == "x" {
			fmt.Println("Bye... ")
			break
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}

		numbers = append(numbers, num)

		sort.Ints(numbers)

		fmt.Println("Sorted slice:", numbers)
	}
}
