package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	Fname, Lname string
}

func main() {
	fmt.Print("Enter filename: ")
	var filename string
	fmt.Scan(&filename)

	file, err := os.Open(filename + ".txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var names []Name

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, " ", 2)

		if len(parts) != 2 {
			fmt.Println("Skipping invalid line:", line)
			continue
		}

		name := Name{
			Fname: parts[0],
			Lname: parts[1],
		}
		names = append(names, name)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Names from file:")
	for _, name := range names {
		fmt.Printf(" [%s %s]", name.Fname, name.Lname)
	}
}
