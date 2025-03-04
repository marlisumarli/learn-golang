package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hello from anonymous function!")
	}()

	func(name string) {
		fmt.Println("Hello,", name)
	}("Golang")

	// Variable func
	add := func(a, b int) int {
		return a + b
	}

	result := add(3, 5)
	fmt.Println("Hasil penjumlahan:", result)
}
