package main

import "fmt"

func main() {
	var n32 = 32768
	var n64 = int64(n32)
	var n16 = int16(n64)

	fmt.Println(n32)
	fmt.Println(n64)
	fmt.Println(n16)

	var name = "Golang"
	var ASCII = name[0]
	var sASCII = string(ASCII)

	fmt.Println(name)
	fmt.Println(ASCII)
	fmt.Println(sASCII)
}
