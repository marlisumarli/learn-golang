package main

import "fmt"

func random() any {
	return true
}

func main() {
	var result any = random()

	// var resultInt = result.(int)
	// fmt.Println(resultInt)

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)

	case int:
		fmt.Println("Int", value)

	default:
		fmt.Println("Unkown", value)
	}
}
