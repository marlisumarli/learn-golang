package main

import "fmt"

func logging() {
	fmt.Println("Funciton executed")
}

func runApp() {
	defer logging()
	fmt.Println("Running")
}

func main() {
	runApp()
}
