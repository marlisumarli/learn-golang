package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	}
	return map[string]string{
		"name": name,
	}
}

func main() {
	data := NewMap("Golang")

	if data == nil {
		fmt.Println("Data masih kosong")
	} else {
		fmt.Println(data["name"])
	}
}
