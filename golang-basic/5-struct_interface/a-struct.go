package main

import "fmt"

// struct ibarat class
type Car struct {
	Brand, Type string
	HP, Seet    int
}

// Membuat method menggunakan () diawal
func (car Car) upgrade(number int) {
	car.HP += number
	fmt.Println("Car upgraded +", car.HP, "HP")
}

func main() {
	car := Car{
		Brand: "Honda",
		Type:  "HRV",
		Seet:  4,
		HP:    10,
	}
	car.upgrade(10)
	fmt.Println(car)
}
