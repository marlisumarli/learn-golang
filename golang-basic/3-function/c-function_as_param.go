package main

import "fmt"

// Operation Lebih mudah dibaca
type Operation func(a int, b int) int

// Traditional
//func processNumbers(a int, b int, operation func(int, int) int) int {
//	return operation(a, b)
//}

// Simple
func processNumbers(a int, b int, operation Operation) int {
	return operation(a, b)
}

func main() {
	add := func(x, y int) int {
		return x + y
	}

	multiply := func(x, y int) int {
		return x * y
	}

	result1 := processNumbers(5, 3, add)
	result2 := processNumbers(5, 3, multiply)

	fmt.Println("Hasil Penjumlahan:", result1)
	fmt.Println("Hasil Perkalian:", result2)
}
