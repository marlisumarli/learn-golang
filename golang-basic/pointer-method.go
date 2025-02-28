package main

import "fmt"

type Man struct {
	Name string
}

func (m *Man) Married() {
	m.Name = "Mr. " + m.Name
}

func main() {
	man := Man{Name: "Koding"}
	man.Married()

	fmt.Println(man)
}
