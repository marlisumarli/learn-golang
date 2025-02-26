package main

import "fmt"

func getHello(name string) string {
	return "Hello " + name
}

func multiHello() (string, string) {
	return "Belajar", "Golang"
}

func main() {
	name := "Golang"
	fmt.Println(getHello(name))

	val1, val2 := multiHello()
	fmt.Println(val1)
	fmt.Println(val2)

	belajar, _ := multiHello() // Bisa hiraukan menggunakan _
	fmt.Println(belajar)
}
