package main

import "fmt"

type Address struct {
	City, Province, Country string
}

// Pointer function
func changeDefaultCity(address *Address) {
	address.City = "Default"
}

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indo"}
	fmt.Println(address1)

	// Pointer mengacu ke address1 menggunakan &
	address2 := &address1

	address2.City = "Jakarta"
	fmt.Println(address1)
	fmt.Println(address2)

	// Ini akan mengubah semua data yang menggunakan address
	*address2 = Address{"Los Santos", "Cikarang", "Batam"}

	fmt.Println(address1)
	fmt.Println(address2)

	// Menggunakan new
	addressNew1 := new(Address)
	addressNew2 := addressNew1

	addressNew2.City = "JKT"

	fmt.Println(addressNew1)
	fmt.Println(addressNew2)

	addressDefault := new(Address)
	changeDefaultCity(addressDefault)

	fmt.Println(addressDefault)
}
