package main

import "fmt"

type Address struct {
	City, Country string
}

// Tidak berubah
//func changeCountry(a Address) {
//	a.Country = "Indonesia"
//}

func changeCountry(a *Address) {
	a.Country = "Indonesia"
}

func main() {
	//address := Address{} // Tidak akan berubah
	address := Address{}
	changeCountry(&address)

	fmt.Println(address)
}
