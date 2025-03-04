package main

import "fmt"

func getHello(name string) string {
	return "Hello " + name
}

// Multiple return
func multiHello() (string, int) {
	return "Belajar", 123
}

func main() {
	name := "Golang"
	fmt.Println(getHello(name))

	val1, val2 := multiHello()
	fmt.Print("Multi return: (", val1, ",", val2, ")\n")

	belajar, _ := multiHello() // Bisa hiraukan menggunakan _
	fmt.Println(belajar)
}
