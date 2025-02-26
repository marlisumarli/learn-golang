package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(name HasName) {
	fmt.Println("Hello", name.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type World struct {
	Name string
}

func (world World) GetName() string {
	return world.Name
}

func main() {
	person := Person{
		Name: "Golang",
	}
	SayHello(person)

	world := World{
		Name: "World",
	}

	SayHello(world)
}
