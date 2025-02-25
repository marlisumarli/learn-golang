package main

import (
	"fmt"
)

func main() {
	var nameOpt string // Optional keyword string
	nameOpt = "Opt var"
	fmt.Println(nameOpt)

	var nameInl = "Inline var" // Inline variable
	fmt.Println(nameInl)

	nameSpl := "Simplefield var" // Simplefield variable
	fmt.Println(nameSpl)

	// Multiline variable
	var (
		first  = "Belajar"
		middle = "Bahasa"
		last   = "Golang"
	)
	fmt.Println(first)
	fmt.Println(middle)
	fmt.Println(last)

	// Constant inline variable
	const version string = "1.0.0"
	fmt.Println(version)

	// Constant multiple variable
	const (
		v    = "1.0.0"
		lang = "go"
		sys  = "Mac"
	)
	fmt.Println(v)
	fmt.Println(lang)
	fmt.Println(sys)
}
