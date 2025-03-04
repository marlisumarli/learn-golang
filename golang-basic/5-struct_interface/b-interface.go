package main

import "fmt"

type Speaker interface {
	Speak() string
}

// Struct yang mengimplementasikan interface
type Person struct {
	Name string
}

// Implementasi Speak()
func (p Person) Speak() string {
	return "Hello, my name is " + p.Name
}

type Cat struct {
	Breed string
}

func (d Cat) Speak() string {
	return "Miaww! I am a " + d.Breed
}

// Fungsi yang menerima parameter bertipe interface
func introduce(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	person := Person{Name: "Alice"}
	dog := Cat{Breed: "Golden Retriever"}

	introduce(person)
	introduce(dog)
}
