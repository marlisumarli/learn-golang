package main

import "fmt"

func main() {

	// Menggunakan variable
	counter := 1

	fmt.Println("Mulai")
	for counter <= 10 {
		fmt.Println("Perulangan ke", counter)
		counter++
	}
	fmt.Println("Selesai")

	fmt.Println("Mulai")
	// Bisa seperti ini
	for i := 1; i <= 10; i++ {
		fmt.Println("Perulangan ke", i)
	}
	fmt.Println("Selesai")

	// Perulangan array
	cars := []string{"BMW", "Honda", "BYD"}
	for i := 0; i < len(cars); i++ {
		fmt.Println(cars[i])
	}

	// Bisa seperti ini
	for index, data := range cars {
		fmt.Println("Index", index, `"Value"`, data)
	}

}
