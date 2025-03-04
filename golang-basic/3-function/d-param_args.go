package main

import "fmt"

// Menggunakan literal (...)
func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	fmt.Println(sumAll(10, 10, 10, 10))

	// Bisa seperti ini
	numbers := []int{10, 10, 10}
	fmt.Println(sumAll(numbers...)) // Jangan lupa beri (...) ketika ingin mengirim
}
