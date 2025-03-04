package main

import "fmt"

func main() {
	person := map[string]string{
		"name": "Marleess",
		"age":  "777",
	}

	fmt.Println(person["name"])

	car := make(map[string]string) // Bisa seperti ini
	car["Brand"] = "Honda"
	car["Type"] = "CRV"
	car["Seet"] = "4"
	car["Pedals"] = "No"

	fmt.Println(car)

	// Hapus data
	delete(car, "Pedals")

	fmt.Println(car)
}
