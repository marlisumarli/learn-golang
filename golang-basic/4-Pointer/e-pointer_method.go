package main

import "fmt"

type Car struct {
	Speed int
}

func (m *Car) Up() {
	m.Speed = m.Speed + 1
}

func main() {
	car := Car{Speed: 100}
	car.Up()

	fmt.Println(car)
}
